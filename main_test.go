package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestReadAndDeletelineV2(t *testing.T) {
	//ReadAndDeletelineV2("/Users/backend001/go/src/goshell/test.proto")

	//mongodbId := "650e947a55c6a87840ac518b"
	//_, err := primitive.ObjectIDFromHex(mongodbId)
	//fmt.Println(err)
	//
	//fmt.Println(time.Unix(1698163200, 0).Format("15:04"))
	//Design()

	//SSHWrite()
	StringFunction()
}

type Response struct {
	Value Values `json:"value"`
}

type Values struct {
	List []ListData `json:"list"`
}

type ListData struct {
	LotteryDrawResult string `json:"lotteryDrawResult"`
}

func Design() {
	url := "https://webapi.sporttery.cn/gateway/lottery/getHistoryPageListV1.qry?gameNo=85&provinceId=0&pageSize=2499&isVerify=1&pageNo=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var resp Response
	resp.Value.List = make([]ListData, 0, 2499)

	_ = json.Unmarshal(body, &resp)

	fileName := "example.txt"
	// 使用os.Create打开文件（如果不存在则创建）
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()

	// 创建一个新的写入器
	writer := bufio.NewWriter(file)

	for _, v := range resp.Value.List {
		// 写入字符串到文件
		_, err = writer.WriteString(fmt.Sprintf("%s\n", v.LotteryDrawResult))
		if err != nil {
			fmt.Println("无法写入文件:", err)
			return
		}
	}
	// 确保所有缓存的操作已应用到底层写入器
	err = writer.Flush()
	if err != nil {
		fmt.Println("无法刷新写入器:", err)
		return
	}

	fmt.Println("文件写入成功!")
}

func SSHWrite() {
	// 解析私钥
	signer, err := ssh.ParsePrivateKey([]byte("-----BEGIN OPENSSH PRIVATE KEY-----\n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\n-----END OPENSSH PRIVATE KEY-----"))
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}
	// 设置 SSH 连接的配置
	config := &ssh.ClientConfig{
		User: "root", // 替换为你的用户名
		Auth: []ssh.AuthMethod{
			//ssh.Password("msg3Ns24LsABC"), // 替换为你的密码
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 不检查服务器的密钥，对于安全性要求高的应用不建议这样做
	}

	// 连接到远程服务器
	client, err := ssh.Dial("tcp", "49.233.12.106:22", config) // 替换为你的服务器地址和端口
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	// 创建一个会话
	session, err := client.NewSession()
	if err != nil {
		client.Close()
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// 执行远程命令
	//var b []byte
	//b, err = session.Output("ls -l") // 替换为你想执行的命令
	//if err != nil {
	//	log.Fatalf("Failed to run: %s", err)
	//}
	//fmt.Println(string(b))

	err = session.Run("cd /data;mkdir testFolder")
	if err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
	// 关闭连接
	client.Close()
}

func StringFunction() {
	waitReplace := "{\n  \"type\": \"object\",\n  \"properties\": {\n    \"organizationId\": {\n      \"type\": \"string\",\n      \"description\": \" 机构id\"\n    },\n    \"moduleId\": {\n      \"type\": \"string\",\n      \"description\": \" 模块id 模块id为空表示新增模块(自定义)\"\n    },\n    \"moduleName\": {\n      \"type\": \"string\",\n      \"description\": \" 模块名称\"\n    },\n    \"moduleWebUrl\": {\n      \"type\": \"string\",\n      \"description\": \" 模块 '更多'链接\"\n    },\n    \"contentList\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"object\",\n        \"properties\": {\n          \"contentId\": {\n            \"type\": \"string\",\n            \"description\": \" 内容id为空表示 新增数据 不为空表示 修改对应id数据\"\n          },\n          \"title\": {\n            \"type\": \"string\",\n            \"description\": \" 标题\"\n          },\n          \"description\": {\n            \"type\": \"string\",\n            \"description\": \" 描述\"\n          },\n          \"imageUrl\": {\n            \"type\": \"string\",\n            \"description\": \" 图片地址\"\n          },\n          \"webUrl\": {\n            \"type\": \"string\",\n            \"description\": \" 链接地址\"\n          },\n          \"subject\": {\n            \"type\": \"string\",\n            \"description\": \" 科目名称\"\n          },\n          \"categoryId\": {\n            \"type\": \"integer\",\n            \"format\": \"int64\",\n            \"description\": \" 科目id\"\n          }\n        },\n        \"title\": \"EditModuleContentData\",\n        \"$$ref\": \"#/definitions/EditModuleContentData\"\n      },\n      \"description\": \" 模块对应的内容\"\n    }\n  },\n  \"title\": \"EditModuleContentReq\",\n  \"required\": [\n    \"organizationId\",\n    \"contentList\"\n  ],\n  \"$$ref\": \"#/definitions/EditModuleContentReq\"\n}"
	fmt.Println(strings.ReplaceAll(waitReplace, "\\n", ""))
}
