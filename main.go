package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func main() {
	var path string
	var service string
	flag.StringVar(&path, "path", "", "")
	flag.StringVar(&service, "service", "", "")
	green := color.New(color.FgHiGreen)
	red := color.New(color.FgRed)
	flag.Parse()
	if path != "" && service != "" {
		//ReadAndDeleteline(path, service)
		ReadAndDeletelineV2(path)
		_, _ = green.Println("Done!")
		return
	}
	_, _ = red.Println("FAIL!")
}

func ReadAndDeleteline(Path string, RpcService string) {

	// 打开文件
	oldFileName := Path
	file, err := os.Open(Path)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()
	newFileName := fmt.Sprintf("%s.temp", Path)
	newFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("无法创建新文件:", err)
		return
	}
	defer newFile.Close()

	// 创建一个新的Scanner对象，用于逐行读取文件
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(newFile)
	CacheService := make(map[string]string, 0)
	CacheServiceList := make([]string, 0)
	var lines []string // 用于存储新文件的内容

	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "message") {
			newline := strings.Replace(line, "{", "", -1) // 去掉符号 {
			newline = strings.Replace(newline, "message", "", 1)
			newline = strings.TrimSpace(newline)
			if strings.Contains(newline, "Req") {
				lineList := strings.Split(newline, "Req")
				if len(lineList) > 0 {
					serviceName := strings.TrimSpace(lineList[0])
					_, ok := CacheService[serviceName]
					if !ok {
						remark := serviceName
						if len(lineList) >= 2 && strings.Contains(lineList[1], "//") {
							remarkList := strings.Split(lineList[1], "{")
							if len(remarkList) >= 1 {
								remark = remarkList[0]
							}
						}
						CacheService[serviceName] = remark
						CacheServiceList = append(CacheServiceList, serviceName)

					}
				}
			}
		}
		line = TrimMorePrefix(line, []string{" ", "\t"})
		// !strings.HasPrefix(line, "service") && !strings.HasPrefix(line, "rpc") && !strings.HasPrefix(line, "    rpc") && !strings.HasPrefix(line, "\trpc")
		if !StringsHasPrefixs(line, []string{"service", "rpc"}) {
			//_, _ = newFile.WriteString(line + "\n")
			if !StringsHasPrefixs(line, []string{"message", "syntax", "package", "option", "}", "import"}) {
				line = "\t" + line
			}
			lines = append(lines, line)
		}
	}

	// 检查是否出现了扫描错误
	if err := scanner.Err(); err != nil {
		fmt.Println("扫描文件时出现错误:", err)
	}
	// 写入新文件的内容，不包括最后一行
	for i, line := range lines {
		if i < len(lines)-1 {
			_, _ = newFile.WriteString(line + "\n")
		}
	}

	_, _ = newFile.WriteString(fmt.Sprintf("service %s {", RpcService) + "\n")
	for _, v := range CacheServiceList {
		res := CacheService[v]
		var remark string
		if v != res {
			remark = res
		}
		_, _ = newFile.WriteString(fmt.Sprintf("\trpc %s (%sReq) returns (%sResp);%s", v, v, v, remark) + "\n")
	}
	_, _ = newFile.WriteString("}" + "\n")
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

// 是否存在subSlice中任意一个字符串前缀
func StringsHasPrefixs(s string, subSlice []string) bool {
	for _, v := range subSlice {
		if strings.HasPrefix(s, v) {
			return true
		}
	}
	return false
}

// 去除多个前缀
func TrimMorePrefix(s string, prefixSlice []string) string {
	for _, v := range prefixSlice {
		s = strings.TrimPrefix(s, v)
	}
	return s
}

func ReadAndDeletelineV2(Path string) {
	// 打开文件
	oldFileName := Path
	file, err := os.Open(Path)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()
	newFileName := fmt.Sprintf("%s.temp", Path)
	newFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("无法创建新文件:", err)
		return
	}
	defer newFile.Close()

	// 创建一个新的Scanner对象，用于逐行读取文件
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(newFile)
	CacheService := make(map[string]string, 0)
	CacheServiceName := make(map[string]string, 0)
	CacheMapServiceList := make(map[string][]string, 0)
	var lines []string // 用于存储新文件的内容

	// 逐行读取文件内容,第一次读取保存所有service方法名称
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "service") {
			newline := strings.Replace(line, "{", "", -1) // 去掉符号 {
			newline = strings.Replace(newline, "service", "", 1)
			newline = strings.TrimSpace(newline)
			CacheServiceName[newline] = newline
		}
	}
	// 重新定位文件指针到文件开头
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("重新定位文件指针时出错:", err)
		return
	}
	scanner2 := bufio.NewScanner(file)
	// 逐行读取文件内容
	for scanner2.Scan() {
		line := scanner2.Text()
		//fmt.Println(line)
		if strings.HasPrefix(line, "message") {
			newline := strings.Replace(line, "{", "", -1) // 去掉符号 {
			newline = strings.Replace(newline, "message", "", 1)
			newline = strings.TrimSpace(newline)
			if strings.Contains(newline, "Req") {
				lineList := strings.Split(newline, "Req")
				if len(lineList) > 0 {
					serviceName := strings.TrimSpace(lineList[0])
					_, ok := CacheService[serviceName]
					if !ok {
						remark := serviceName
						if len(lineList) >= 2 && strings.Contains(lineList[1], "//") {
							remarkList := strings.Split(lineList[1], "{")
							if len(remarkList) >= 1 {
								remark = remarkList[0]
							}
						}
						CacheService[serviceName] = remark
						IsHas, Name := StringsHasPrefixInMap(serviceName, CacheServiceName)

						if IsHas {
							serviceList, OK := CacheMapServiceList[Name]
							if !OK {
								serviceList = make([]string, 0)
							}
							serviceList = append(serviceList, serviceName)
							CacheMapServiceList[Name] = serviceList
						}
					}
				}
			}
		}
		fmt.Println(line)
		line = TrimMorePrefix(line, []string{" ", "\t"})
		// !strings.HasPrefix(line, "service") && !strings.HasPrefix(line, "rpc") && !strings.HasPrefix(line, "    rpc") && !strings.HasPrefix(line, "\trpc")
		if !StringsHasPrefixs(line, []string{"rpc"}) {
			//_, _ = newFile.WriteString(line + "\n")
			if !StringsHasPrefixs(line, []string{"message", "syntax", "package", "option", "}", "service", "import"}) {
				line = "\t" + line
			}
			lines = append(lines, line)
		}
	}

	// 检查是否出现了扫描错误
	if err := scanner.Err(); err != nil {
		fmt.Println("扫描文件时出现错误:", err)
	}

	// 写入新文件的内容
	for _, line := range lines {
		_, _ = newFile.WriteString(line + "\n")
		if strings.HasPrefix(line, "service") {
			newline := strings.Replace(line, "{", "", -1) // 去掉符号 {
			newline = strings.Replace(newline, "service", "", 1)
			newline = strings.TrimSpace(newline)
			//fmt.Println(newline)
			ServiceList, ok := CacheMapServiceList[newline]
			if ok {
				for _, v := range ServiceList {
					res := CacheService[v]
					var remark string
					if v != res {
						remark = res
					}
					_, _ = newFile.WriteString(fmt.Sprintf("\trpc %s (%sReq) returns (%sResp);%s", v, v, v, remark) + "\n")
				}
			}
		}
	}
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

func StringsHasPrefixInMap(s string, mapdata map[string]string) (bool, string) {
	for _, v := range mapdata {
		if strings.HasPrefix(s, v) {
			return true, v
		}
	}
	return false, ""
}
