package newXsql

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"regexp"
	"strings"
)

type SqlTrans struct {
	FilePath       string // 文件路径
	DataSourceName string // 数据库连接字符串
}

func NewSqlTrans(filePath, dataSourceName string) ISqlTrans {
	return &SqlTrans{
		FilePath:       filePath,
		DataSourceName: dataSourceName,
	}
}

func (s *SqlTrans) SqlTransGoStructs() {
	// 打开 SQL 文件
	file, err := os.Open(s.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 读取 SQL 文件内容
	var createTableSQL string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		createTableSQL += scanner.Text() + " "
	}

	// 解析 SQL 语句，提取表名和列定义
	tableStart := strings.Index(createTableSQL, "CREATE TABLE")
	if tableStart == -1 {
		log.Fatal("No CREATE TABLE statement found in SQL file.")
	}
	createTableSQL = createTableSQL[tableStart:]

	tableEnd := strings.Index(createTableSQL, "(")
	if tableEnd == -1 {
		log.Fatal("Malformed CREATE TABLE statement in SQL file.")
	}
	tableName := strings.TrimSpace(createTableSQL[len("CREATE TABLE"):tableEnd])
	tableColumns := createTableSQL[tableEnd+1 : len(createTableSQL)-1]

	// 打印表名
	fmt.Println("Table Name:", tableName)

	// 打印生成的 Go 结构体代码
	fmt.Println("Generated Go Struct:")

	// 根据列定义生成 Go 结构体字段
	columnDefs := strings.Split(tableColumns, ",")
	fmt.Println("type", tableName, "struct {")
	for _, columnDef := range columnDefs {
		comIndex := strings.Index(columnDef, "COMMENT")
		columnParts := strings.Fields(columnDef)
		if len(columnParts) >= 2 {
			columnName := strings.Trim(columnParts[0], "`")
			columnType := convertColumnType(columnParts[1])
			var comment string
			if comIndex >= 0 {
				comment = columnDef[comIndex+8:]
				comment = strings.ReplaceAll(comment, "'", "")
				comment = fmt.Sprintf(" // %s", comment)
			}
			fmt.Printf("    %s %s `json:\"%s\" db:\"%s\"` %s\n", toCamelCase(columnName), columnType, columnName, columnName, comment)
		}
	}
	fmt.Println("}")
}

func (s *SqlTrans) SqlTransProtoMessages() {
	// 打开 SQL 文件
	file, err := os.Open(s.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 读取 SQL 文件内容
	var createTableSQL string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		createTableSQL += scanner.Text() + " "
	}

	// 解析 SQL 语句，提取表名和列定义
	tableStart := strings.Index(createTableSQL, "CREATE TABLE")
	if tableStart == -1 {
		log.Fatal("No CREATE TABLE statement found in SQL file.")
	}
	createTableSQL = createTableSQL[tableStart:]

	tableEnd := strings.Index(createTableSQL, "(")
	if tableEnd == -1 {
		log.Fatal("Malformed CREATE TABLE statement in SQL file.")
	}
	tableName := strings.TrimSpace(createTableSQL[len("CREATE TABLE"):tableEnd])
	tableColumns := createTableSQL[tableEnd+1 : len(createTableSQL)-1]

	// 打印表名
	fmt.Println("Table Name:", tableName)

	// 打印生成的 Go 结构体代码
	fmt.Println("Generated Go Struct:")

	// 根据列定义生成 Go 结构体字段
	columnDefs := strings.Split(tableColumns, ",")
	fmt.Println("message", strings.Trim(tableName, "`"), " {")
	index := 1
	for _, columnDef := range columnDefs {
		comIndex := strings.Index(columnDef, "COMMENT")
		columnParts := strings.Fields(columnDef)
		if len(columnParts) >= 2 {
			columnName := strings.Trim(columnParts[0], "`")
			columnType := convertColumnType(columnParts[1])
			//fmt.Printf("    %s %s `json:\"%s\" db:\"%s\"`\n", toCamelCase(columnName), columnType, columnName, columnName)
			if columnType == "time.Time" {
				columnType = "google.protobuf.Timestamp"
			}
			var comment string
			if comIndex >= 0 {
				comment = columnDef[comIndex+8:]
				comment = strings.ReplaceAll(comment, "'", "")
				comment = fmt.Sprintf(" // %s", comment)
			}
			fmt.Printf(" %s %s = %d;%s\n", columnType, columnName, index, comment)
			index += 1
		}
	}
	fmt.Println("}")
}

const (
	tableCreatePrefix  = "CREATE TABLE"
	principalKeyPrefix = "PRIMARY KEY"
	engineInnoDB       = ") ENGINE=InnoDB"
)

var re = regexp.MustCompile("`([^`]+)`") // 移出循环，作为全局变量

func (s *SqlTrans) SqlTableAddColumns() {
	db, err := sql.Open("mysql", s.DataSourceName)
	if err != nil {
		fmt.Println("Failed to open database connection:", err)
		return
	}
	defer db.Close()

	file, err := os.Open(s.FilePath)
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tableName string
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if !stringsHasPrefixs(line, []string{tableCreatePrefix, principalKeyPrefix, engineInnoDB}) {
			processLine(db, line, tableName)
		} else if strings.HasPrefix(line, tableCreatePrefix) {
			matches := re.FindStringSubmatch(line)
			tableName = matches[1]
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error occurred while scanning file:", err)
	}
}

func processLine(db *sql.DB, line, tableName string) {
	matches := re.FindStringSubmatch(line)
	field := matches[1]
	_, comment, _ := strings.Cut(line, field)
	comment = strings.TrimRight(comment, ",")
	query := fmt.Sprintf("SHOW COLUMNS FROM %s LIKE '%s'", tableName, field)
	//fmt.Println(query)
	// 使用参数化查询以避免SQL注入
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	var data DBData
	err = stmt.QueryRow().Scan(&data.Fields, &data.Types, &data.Nulls, &data.Keys, &data.Defaults, &data.Extras)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return // 确保在错误发生时优雅地返回
	}
	if data.Fields == "" { // 检查列是否不存在
		res := fmt.Sprintf(`ALTER TABLE %s ADD COLUMN %s %s;`, tableName, field, comment)
		fmt.Println(res)
	}
}

type DBData struct {
	Fields   string      `db:"Field"`
	Types    string      `db:"Type"`
	Nulls    string      `db:"Null"`
	Keys     interface{} `db:"Key"`
	Defaults interface{} `db:"Default"`
	Extras   string      `db:"Extra"`
}

func stringsHasPrefixs(s string, subSlice []string) bool {
	for _, v := range subSlice {
		if strings.HasPrefix(s, v) {
			return true
		}
	}
	return false
}

// 将 SQL 列类型转换为 Go 类型
func convertColumnType(sqlType string) string {
	if strings.Contains(sqlType, "(") {
		sqlTypeList := strings.Split(sqlType, "(")
		sqlType = sqlTypeList[0]
	}
	switch sqlType {
	case "int", "smallint", "bigint":
		return "int64"
	case "tinyint":
		return "int"
	case "varchar", "text", "char":
		return "string"
	case "date", "datetime", "timestamp":
		return "time.Time"
	default:
		return "interface{}"
	}
}

// 转换字符串为驼峰命名
func toCamelCase(input string) string {
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == '_' || r == ' ' || r == '-'
	})

	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}
