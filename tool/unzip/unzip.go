package unzip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var CdnCosBucket = map[string]string{
	"https://cdnninebeats.wedomusic.cn":                             "package-9beats-1252905615",
	"http://cdnninebeats.wedomusic.cn":                              "package-9beats-1252905615",
	"https://duolaixuecdn.wedomusic.cn":                             "package-duolaixue-1252905615",
	"http://duolaixuecdn.wedomusic.cn":                              "package-duolaixue-1252905615",
	"https://package-9beats-1252905615.cos.ap-beijing.myqcloud.com": "package-9beats-1252905615",
	"http://package-9beats-1252905615.cos.ap-beijing.myqcloud.com":  "package-9beats-1252905615",
	"https://cdn9beatsold.wedomusic.cn":                             "9beats-old-1252905615",
	"http://cloud.cdn.wedomusic.cn":                                 "dlx-cloud-disk-1252905615",
	"https://cloud.cdn.wedomusic.cn":                                "dlx-cloud-disk-1252905615",
	//"http://cdnshangda.wedomusic.cn":                                "package-shangda-1252905615",
	//"https://cdnshangda.wedomusic.cn":                               "package-shangda-1252905615",
	//"https://cdnyuequ.wedomusic.cn":                                 "package-yuequ-1252905615",
	//"http://cdnyuequ.wedomusic.cn":                                  "package-yuequ-1252905615",
	"https://cdn.cloud-dev.yuetumusic.cn": "aiopc-cloud-dev-1252905615",
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		// 构建解压后的文件路径
		path := filepath.Join(dest, f.Name)
		if strings.HasPrefix(f.Name, "__MACOSX/") {
			continue
		}
		// 检查文件是否是文件夹
		if f.FileInfo().IsDir() {
			// 如果是文件夹，创建对应的文件夹
			os.MkdirAll(path, f.Mode())
			continue
		}

		// 如果是文件，创建对应的文件
		if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func deleteFilesAndFolder() {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前目录:", err)
		return
	}

	// 获取当前目录下的所有文件和文件夹
	files, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Println("无法读取当前目录:", err)
		return
	}

	// 遍历当前目录下的所有文件和文件夹
	for _, file := range files {
		// 删除.zip文件
		if !file.IsDir() && filepath.Ext(file.Name()) == ".zip" {
			err := os.Remove(file.Name())
			if err != nil {
				fmt.Println("删除文件失败:", err)
			} else {
				fmt.Println("删除文件成功:", file.Name())
			}
		}

		// 删除名为unzipped的文件夹
		if file.IsDir() && file.Name() == "unzipped" {
			err := os.RemoveAll(file.Name())
			if err != nil {
				fmt.Println("删除文件夹失败:", err)
			} else {
				fmt.Println("删除文件夹成功:", file.Name())
			}
		}
	}
}

func findImageFileWith01Suffix() (string, error) {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("无法获取当前目录：%v", err)
	}

	// 构建unzipped目录的路径
	unzippedDir := filepath.Join(currentDir, "unzipped")

	// 用于保存找到的文件全名
	var foundFilePath string

	// 遍历unzipped目录下的所有文件
	err = filepath.Walk(unzippedDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查文件是否直接位于unzipped目录下
		dir := filepath.Dir(path)
		if dir == unzippedDir {
			// 检查文件是否为图片文件（后缀名为.jpg、.png、.gif等）
			isImage, ext := isImageFile(path)
			if !info.IsDir() && isImage {
				// 检查文件名是否以1结尾
				fmt.Println("Name:", info.Name())
				if strings.HasSuffix(info.Name(), fmt.Sprintf("1%s", ext)) {
					foundFilePath = path
					return fmt.Errorf("找到符合条件的图片文件：%s", path)
				}
			}
		}
		return nil
	})

	if err != nil && foundFilePath == "" {
		return "", err
	}

	return foundFilePath, nil
}

func isImageFile(filename string) (bool, string) {
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
		return true, ext
	}
	return false, ""
}

func findSmallestImageFile() (string, error) {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("无法获取当前目录：%v", err)
	}

	// 构建unzipped目录的路径
	unzippedDir := filepath.Join(currentDir, "unzipped")

	// 用于保存找到的文件全名
	var smallestImage string

	// 读取文件夹内容
	files, err := ioutil.ReadDir(unzippedDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return "", err
	}

	var imageFiles []string
	for _, file := range files {
		if !file.IsDir() {
			isImage, _ := isImageFile(file.Name())
			if isImage {
				imageFiles = append(imageFiles, file.Name())
			}
		}
	}
	// 自定义排序函数，按文件名结尾的数字排序
	sort.Slice(imageFiles, func(i, j int) bool {
		return extractNumber(imageFiles[i]) < extractNumber(imageFiles[j])
	})

	// 输出排序后的第一个图片文件名
	if len(imageFiles) > 0 {
		smallestImage = imageFiles[0]
		return smallestImage, nil
	} else {
		err = errors.New("没有图片文件")
		return "", nil
	}
}

func extractNumber(filename string) int {
	parts := strings.Split(filename, "_")
	if len(parts) > 1 {
		numberPart := strings.TrimSuffix(parts[len(parts)-1], filepath.Ext(parts[len(parts)-1]))
		number, err := strconv.Atoi(numberPart)
		if err == nil {
			return number
		}
	}
	return -1
}
