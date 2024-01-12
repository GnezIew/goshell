package imageToPDF

import (
	"fmt"
	"testing"
)

func TestImageToPDF(t *testing.T) {
	//dirPath := "/Users/backend001/go/src/goshell/tool/imageToPDF/image"
	//ImageToPDF(dirPath)
	//EndTime := "2023-12-21"
	//EndTimeT, _ := time.ParseInLocation("2006-01-02", EndTime, time.Local)
	//NowTimeT, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	//fmt.Println(EndTimeT.Sub(NowTimeT).Hours() / 24)
	serviceAmount := 8800

	fmt.Println(int64(float64(serviceAmount) / 30 * float64(28)))
	fmt.Println(8819 / 30)
	fmt.Println(8819 % 30)
}
