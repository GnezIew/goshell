package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type GlobalVal struct {
	FilePath       string `json:"filePath"`
	DataSourceName string `json:"dataSourceName"`
}

// 定一个全局的对外的GlobalObj
var GlobalObject *GlobalVal

func (g *GlobalVal) ReloadConf() {
	data, err := os.ReadFile("conf/xsql.json")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		fmt.Println("unmarshal err:", err)
		return
	}
}

func init() {
	GlobalObject = new(GlobalVal)
	GlobalObject.ReloadConf()
}
