package unzip

import (
	"fmt"
	"testing"
)

func TestUnZip(t *testing.T) {
	//Path, _ := os.Getwd()
	//oss.DownLoad("package-9beats-1252905615", "textbook_zip/1530517144332.zip", fmt.Sprintf("%s/%s", Path, "1530517144332.zip"))

	//zipFile := "1530517144332.zip" // 要解压的.zip文件
	//destFolder := "unzipped"       // 解压后的目标文件夹
	//
	//err := Unzip(zipFile, destFolder)
	//if err != nil {
	//	fmt.Println("解压失败:", err)
	//} else {
	//	fmt.Println("解压成功")
	//}
	//// 查询unzipped目录下是否有01结尾的.png或.jpg的文件
	//fileName, err := findImageFileWith01Suffix()
	//fmt.Println(fileName, err)
	//// 上传图片
	//url, err := oss.PutFromFile("package-9beats-1252905615", fmt.Sprintf("%s_%s", "textbook_zip/1530517144332", "tupian_0001.png"), fileName)
	//fmt.Println(url, err)
	//deleteFilesAndFolder()

	url := UnZipAndUpLoadOss("https://cdn.cloud-dev.yuetumusic.cn/testZip/1555666265787.zip", "https://cdnninebeats.wedomusic.cn/sibelius_json/1555666294_sibei.json")
	fmt.Println(url)
}
