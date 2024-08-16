package xswagger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type XSwagger struct {
	Swagger             string                 `json:"swagger"`
	Info                interface{}            `json:"info"`
	Schemes             interface{}            `json:"schemes"`
	Consumes            interface{}            `json:"consumes"`
	Produces            interface{}            `json:"produces"`
	Paths               map[string]interface{} `json:"paths"`
	Definitions         interface{}            `json:"definitions"`
	SecurityDefinitions interface{}            `json:"securityDefinitions"`
}

func NewXswagger() IxSwagger {
	return &XSwagger{}
}

func (x *XSwagger) CompareSwaggerJson(oriFilePath string, newFilePath string) string {
	fs, err := os.Open(oriFilePath)
	if err != nil {
		return err.Error()
	}
	defer fs.Close()
	jsonData, err := io.ReadAll(fs)
	if err != nil {
		return err.Error()
	}
	err = json.Unmarshal(jsonData, &x)
	if err != nil {
		return err.Error()
	}

	fs2, err := os.Open(newFilePath)
	if err != nil {
		return err.Error()
	}
	defer fs2.Close()
	jsonData2, err := io.ReadAll(fs2)
	if err != nil {
		return err.Error()
	}
	var newData XSwagger
	err = json.Unmarshal(jsonData2, &newData)
	if err != nil {
		return err.Error()
	}
	differentList := make([]string, 0)
	for k, v := range x.Paths {
		newVal, ok := newData.Paths[k]
		if !ok {
			differentList = append(differentList, k)
		} else {
			vStr := v.(string)
			newValStr := newVal.(string)
			if vStr != newValStr {
				differentList = append(differentList, k)
			}
		}
	}
	fmt.Println(differentList)
	return ""
}
