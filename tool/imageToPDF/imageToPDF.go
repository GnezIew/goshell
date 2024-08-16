package imageToPDF

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type ImageData struct {
	DirPath string
}

func NewImageData(dirPath string) ImageDataInterface {
	return &ImageData{DirPath: dirPath}
}

func (i *ImageData) ImageToPDF() {
	imageList, err := getFolderImageList(i.DirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 对图片进行排序
	sort.Slice(imageList, func(i, j int) bool {
		return extractNumber(imageList[i]) < extractNumber(imageList[j])
	})
	pdf := gofpdf.New("P", "cm", "A4", "") // 初始化
	imageWidthPt := 1334.0 / 72.0          // 将宽度转换为点单位
	imageHeightPt := 750.0 / 72.0          // 将高度转换为点单位
	for _, imageFile := range imageList {
		pdf.AddPageFormat("P", gofpdf.SizeType{
			Wd: imageWidthPt,
			Ht: imageHeightPt,
		})
		// 获取图片地址后缀
		options := gofpdf.ImageOptions{
			ImageType: getImageExt(imageFile),
			ReadDpi:   true,
		}
		pdf.ImageOptions(imageFile, 0, 0, imageWidthPt, imageHeightPt, false, options, 0, "")
	}
	err = pdf.OutputFileAndClose("output.pdf") // 保存为PDF
	if err != nil {
		panic(err)
	}
}

func ImageToPDF(dirPath string) {
	imageList, err := getFolderImageList(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 对图片进行排序
	sort.Slice(imageList, func(i, j int) bool {
		return extractNumber(imageList[i]) < extractNumber(imageList[j])
	})
	pdf := gofpdf.New("P", "cm", "A4", "") // 初始化
	imageWidthPt := 1334.0 / 72.0          // 将宽度转换为点单位
	imageHeightPt := 750.0 / 72.0          // 将高度转换为点单位
	for _, imageFile := range imageList {
		pdf.AddPageFormat("P", gofpdf.SizeType{
			Wd: imageWidthPt,
			Ht: imageHeightPt,
		})
		// 获取图片地址后缀
		options := gofpdf.ImageOptions{
			ImageType: getImageExt(imageFile),
			ReadDpi:   true,
		}
		pdf.ImageOptions(imageFile, 0, 0, imageWidthPt, imageHeightPt, false, options, 0, "")
	}
	err = pdf.OutputFileAndClose("output.pdf") // 保存为PDF
	if err != nil {
		panic(err)
	}
}

// 解析文件夹
func getFolderImageList(dirPath string) ([]string, error) {
	var imageFiles []string
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if isImageFile(file.Name()) {
			imageFiles = append(imageFiles, filepath.Join(dirPath, file.Name()))
		}
	}
	return imageFiles, nil
}

// 是否图片文件
func isImageFile(fileName string) bool {
	extensions := []string{".jpg", ".jpeg", ".png"}
	for _, ext := range extensions {
		if strings.HasSuffix(strings.ToLower(fileName), ext) {
			return true
		}
	}
	return false
}

// 获取图片文件后缀名
func getImageExt(fileName string) string {
	fileNameList := strings.Split(fileName, ".")
	ext := fileNameList[len(fileNameList)-1]
	return strings.ToUpper(ext)
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
