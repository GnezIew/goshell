package storage

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func SqlTransGoStruct() {
	// 打开 SQL 文件
	file, err := os.Open("create_table.sql")
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
		columnParts := strings.Fields(columnDef)
		if len(columnParts) >= 2 {
			columnName := strings.Trim(columnParts[0], "`")
			columnType := convertColumnType(columnParts[1])
			fmt.Printf("    %s %s `json:\"%s\" db:\"%s\"`\n", toCamelCase(columnName), columnType, columnName, columnName)
		}
	}
	fmt.Println("}")
}

func SqlTransProtoMessage() {
	// 打开 SQL 文件
	file, err := os.Open("create_table.sql")
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
		columnParts := strings.Fields(columnDef)
		if len(columnParts) >= 2 {
			columnName := strings.Trim(columnParts[0], "`")
			columnType := convertColumnType(columnParts[1])
			//fmt.Printf("    %s %s `json:\"%s\" db:\"%s\"`\n", toCamelCase(columnName), columnType, columnName, columnName)
			if columnType == "time.Time" {
				columnType = "google.protobuf.Timestamp"
			}
			fmt.Printf(" %s %s = %d;\n", columnType, columnName, index)
			index += 1
		}
	}
	fmt.Println("}")
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

type TransData struct {
	PlaceHolder string
	QueryIds    []interface{}
}

func TransInSql(waitTransData []interface{}, prefixAppend ...interface{}) *TransData {
	if len(waitTransData) == 0 {
		return nil
	}
	transData := &TransData{ // 初始化一个 TransData 结构体并获得指针
		PlaceHolder: "",
		QueryIds:    make([]interface{}, 0),
	}
	if prefixAppend != nil { // 存在需要前缀追加
		transData.QueryIds = append(transData.QueryIds, prefixAppend...)
	}
	for i := 0; i < len(waitTransData); i++ {
		transData.PlaceHolder += "?,"
		transData.QueryIds = append(transData.QueryIds, waitTransData[i])
	}
	transData.PlaceHolder = transData.PlaceHolder[:len(transData.PlaceHolder)-1]
	return transData
}

const dbTag = "db"

type ParamData struct {
	Param     interface{}
	PageStart int64
	PageEnd   int64
}

func TransParamToWhereSql(param ParamData) (string, error) {
	v := reflect.ValueOf(param.Param)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("字段param必须赋值结构体类型")
	}

	typ := v.Type()
	query := ""
	queryIds := make([]interface{}, 0)
	var sortPlaceHolder, sortValue []string

	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		fieldValue := v.Field(i)
		tagv := fi.Tag.Get(dbTag)
		zeroValue := reflect.Zero(fieldValue.Type()).Interface()
		if !reflect.DeepEqual(fieldValue.Interface(), zeroValue) {
			fieldName, method := splitNameAndMethod(tagv)
			switch method {
			case "", "search":
				if method != "search" {
					query += fmt.Sprintf(" `%s` = ? and", fieldName)
				} else {
					query += fmt.Sprintf(" `%s` like %%?%% and", fieldName)
				}
				queryIds = append(queryIds, fieldValue.Interface())
			case "sort":
				sortPlaceHolder = append(sortPlaceHolder, fieldName)
				sortValue = append(sortValue, fieldValue.String())
			case "in":
				transdata := TransInSql(fieldValue.Interface().([]interface{}))
				query += fmt.Sprintf(" `%s` IN(%s) and", fieldName, transdata.PlaceHolder)
				queryIds = append(queryIds, transdata.QueryIds...)
			}
		}
	}
	query = strings.TrimSuffix(query, "and")

	if len(sortPlaceHolder) != 0 && len(sortValue) != 0 && len(sortValue) == len(sortPlaceHolder) {
		for i := 0; i < len(sortPlaceHolder); i++ {
			if i == 0 {
				query += fmt.Sprintf(" order by %s %s", sortPlaceHolder[i], sortValue[i])
			} else {
				query += fmt.Sprintf(", %s %s", sortPlaceHolder[i], sortValue[i])
			}
		}
	}
	if param.PageEnd != 0 {
		query += fmt.Sprintf(" limit %d,%d", param.PageStart, param.PageEnd)
	} else if param.PageStart != 0 {
		query += fmt.Sprintf(" limit %d", param.PageStart)
	}
	fmt.Println(queryIds)
	return query, nil
}

func sqlInByNum(Num int64) string {
	Nums := int(Num)
	var placeHolder string
	for i := 0; i < Nums; i++ {
		placeHolder += "?,"
	}
	placeHolder = placeHolder[:len(placeHolder)-1]
	return placeHolder
}

func splitNameAndMethod(tagv string) (name, method string) {
	splitList := strings.Split(tagv, ",")
	if len(splitList) < 2 {
		return tagv, ""
	}
	return splitList[0], splitList[1]
}
