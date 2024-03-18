package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	Data apiDocData `json:"data"`
}

type apiDocData struct {
	Title        string `json:"title"`
	Path         string `json:"path"`
	ReqBodyOther string `json:"req_body_other"`
}

type reqBodyOther struct {
	Properties map[string]properties `json:"properties"`
}

type properties struct {
	Types       string `json:"type"`
	Description string `json:"description"`
}

func main() {
	var apiUrl string
	flag.StringVar(&apiUrl, "url", "", "")
	green := color.New(color.FgHiGreen)
	red := color.New(color.FgRed)
	blue := color.New(color.FgCyan)
	flag.Parse()
	apiUrlList := strings.Split(apiUrl, "api/")
	if len(apiUrlList) < 2 {
		_, _ = red.Println("URL FAIL!")
	}
	apiId := apiUrlList[1]

	url := fmt.Sprintf("http://152.136.176.149:3000/api/interface/get?id=%s", apiId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		_, _ = red.Println(err)
		return
	}
	req.Header.Add("Cookie", "_yapi_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIzLCJpYXQiOjE3MTA2NjY4NTEsImV4cCI6MTcxMTI3MTY1MX0.9MnOY5jGBREWI7bGC6yXpuJkpk_XJsKL0Kq0nM0h6FM; _yapi_uid=23")

	res, err := client.Do(req)
	if err != nil {
		_, _ = red.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		_, _ = red.Println(err)
		return
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		_, _ = red.Println(err)
	}

	reqBodyOtherStr := strings.ReplaceAll(response.Data.ReqBodyOther, "\n", "")
	var reqData reqBodyOther
	_ = json.Unmarshal([]byte(reqBodyOtherStr), &reqData)
	var param string
	for k, v := range reqData.Properties {
		param += fmt.Sprintf("%s %s类型", k, v.Types)
		if v.Description != "" {
			param += fmt.Sprintf("(%s)", v.Description)
		}
		param += "、"
	}
	var content string
	content = fmt.Sprintf("%s - %s(%s) - 请求参数: %s 返回参数详见文档", apiUrl, response.Data.Title, response.Data.Path, param)
	_, _ = blue.Println(content)
	_, _ = green.Println("done!")
}
