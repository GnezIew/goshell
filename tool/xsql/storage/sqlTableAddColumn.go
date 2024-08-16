package storage

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"regexp"
	"strings"
)

type DBData struct {
	Fields   string      `db:"Field"`
	Types    string      `db:"Type"`
	Nulls    string      `db:"Null"`
	Keys     interface{} `db:"Key"`
	Defaults interface{} `db:"Default"`
	Extras   string      `db:"Extra"`
}

func TableAddColumn(path string) {
	db, err := sql.Open("mysql", "liubei:gsmfts43mgsGSf@tcp(bj-cdb-b9dtj5s4.sql.tencentcdb.com:63752)/learn?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()

	// 创建一个新的Scanner对象，用于逐行读取文件
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("`([^`]+)`")
	var tableName string
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if !stringsHasPrefixs(line, []string{"CREATE TABLE", "PRIMARY KEY", ") ENGINE=InnoDB"}) {
			// 查找匹配的子字符串
			matches := re.FindStringSubmatch(line)
			field := matches[1]
			_, comment, _ := strings.Cut(line, field)
			comment = strings.TrimRight(comment, ",")
			query := fmt.Sprintf("SHOW COLUMNS FROM %s LIKE '%s'", tableName, field)
			//fmt.Println(query)
			DbRes, err := db.Query(query)
			if err != nil {
				fmt.Println(err)
				continue
			}
			data := new(DBData)
			if DbRes.Next() {
				_ = DbRes.Scan(&data.Fields, &data.Types, &data.Nulls, &data.Keys, &data.Defaults, &data.Extras)
			} else {
				res := fmt.Sprintf(`
ALTER TABLE %s
ADD COLUMN %s %s;
			`, tableName, field, comment)
				fmt.Println(res)
			}
		} else if strings.HasPrefix(line, "CREATE TABLE") {
			matches := re.FindStringSubmatch(line)
			tableName = matches[1]
			//fmt.Println(tableName)
		}
	}
}

func stringsHasPrefixs(s string, subSlice []string) bool {
	for _, v := range subSlice {
		if strings.HasPrefix(s, v) {
			return true
		}
	}
	return false
}

const (
	tableCreatePrefix  = "CREATE TABLE"
	principalKeyPrefix = "PRIMARY KEY"
	engineInnoDB       = ") ENGINE=InnoDB"
)

var re = regexp.MustCompile("`([^`]+)`") // 移出循环，作为全局变量

func TableAddColumn2(path string) {
	db, err := sql.Open("mysql", "liubei:gsmfts43mgsGSf@tcp(bj-cdb-b9dtj5s4.sql.tencentcdb.com:63752)/learn?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("Failed to open database connection:", err)
		return
	}
	defer db.Close()

	file, err := os.Open(path)
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
	// 使用参数化查询以避免SQL注入
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	var data DBData
	err = stmt.QueryRow().Scan(&data.Fields, &data.Types, &data.Nulls, &data.Keys, &data.Defaults, &data.Extras)
	if err != nil {
		fmt.Println(err)
		return // 确保在错误发生时优雅地返回
	}
	if data.Fields == "" { // 检查列是否不存在
		res := fmt.Sprintf(`ALTER TABLE %s ADD COLUMN %s %s;`, tableName, field, comment)
		fmt.Println(res)
	}
}
