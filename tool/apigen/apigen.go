package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/GnezIew/goshell/tool/xstring"
	"github.com/fatih/color"
	"os"
	"strings"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "")
	flag.Parse()
	green := color.New(color.FgHiGreen)
	_, _ = green.Println("start!")
	readAndDelete(path)
}

func readAndDelete(path string) {
	fmt.Println(path)
	// 打开文件
	oldFileName := path
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()
	newFileName := fmt.Sprintf("%s.temp", path)
	newFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("无法创建新文件:", err)
		return
	}
	defer newFile.Close()

	// 创建一个新的Scanner对象，用于逐行读取文件
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(newFile)
	cacheServiceMap := make(map[string]string)
	serviceList := make([]string, 0)
	var lines []string // 用于存储新文件的内容
	var group string
	var service string
	for scanner.Scan() {
		line := scanner.Text()
		newline := line
		line = strings.TrimSpace(line) // 去除左右多余的空格
		if !strings.HasPrefix(line, "service") {
			if strings.Contains(line, "Req") {
				line = strings.TrimSpace(line)
				line = strings.ReplaceAll(line, " ", "")
				lineList := strings.Split(line, "Req")
				serviceName := lineList[0]
				cacheServiceMap[serviceName] = ""
				serviceList = append(serviceList, serviceName)
				if len(lineList) >= 2 && strings.Contains(lineList[1], "//") {
					remarkList := strings.Split(lineList[1], "//")
					if len(remarkList) >= 1 {
						cacheServiceMap[serviceName] = remarkList[1]
					}
				}
			} else if strings.Contains(line, "group") {
				line = strings.ReplaceAll(line, " ", "")
				lineList := strings.Split(line, "group:")
				group = lineList[1]
			}
			lines = append(lines, newline)
		} else {
			line = strings.Replace(line, "service", "", -1)
			line = strings.TrimSpace(line)
			lineList := strings.Split(line, "-api")
			service = lineList[0]
			break
		}
	}
	for _, line := range lines {
		_, _ = newFile.WriteString(line + "\n")
	}
	_, _ = newFile.WriteString(fmt.Sprintf("service %s-api {\n", service))
	for _, v := range serviceList {
		_, _ = newFile.WriteString(fmt.Sprintf("\t@doc \"%s\"\n", cacheServiceMap[v]))
		_, _ = newFile.WriteString(fmt.Sprintf("\t@handler %s\n", v))
		_, _ = newFile.WriteString(fmt.Sprintf("\tpost /%s/%s/%s (%s) returns (%s)\n", service, group, xstring.ToLowerFisrt(v), v+"Req", v+"Resp"))
		_, _ = newFile.WriteString("\n")
	}
	_, _ = newFile.WriteString("}")
	// 刷新写入缓冲区并保存新文件
	writer.Flush()

	// 关闭旧文件
	file.Close()

	// 替换旧文件
	err = os.Rename(newFileName, oldFileName)
	if err != nil {
		fmt.Println("替换旧文件时出现错误:", err)
		return
	}
}
