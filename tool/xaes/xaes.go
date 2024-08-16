package xaes

//
//import (
//	"bytes"
//	"crypto/aes"
//	"crypto/cipher"
//	"crypto/rand"
//	"encoding/base64"
//	"fmt"
//	"io"
//	mathRand "math/rand"
//)
//
//func XAes() {
//	// 要加密的数据
//	plaintext := []byte(`{
//		"paperName": "",
//		"problemList": [
//			{
//				"questionProblemId": "4vcic91i-queprm",
//				"questionClassesId": "4v75b5i6-quebankte",
//				"questionClassesName": "TEXT B",
//				"problemDescription": "这是测试拼音数据",
//				"problemDescriptionPinyin": "[\"zhè\",\"shì\",\"cè\",\"shì\",\"pīn\",\"yīn\",\"shù\",\"jù\"]",
//				"problemAudioUrl": "https://cdntestduolaixue.wedomusic.cn/tk/1711520341.mp3",
//				"problemImageUrl": "",
//				"isShowPinyin": 0,
//				"questionSmallProblemList": [
//					{
//						"questionSmallProblemId": "4vcic91s-quesmprm",
//						"title": "1111111111",
//						"titlePinyin": "",
//						"smallProblemType": 5,
//						"smallProblemDescription": "",
//						"smallProblemDescriptionPinyin": "",
//						"smallProblemAnalysis": "223333333333333333",
//						"score": 1,
//						"audioUrl": "",
//						"imageUrl": "",
//						"questionOptionList": [
//							{
//								"problemOptionId": "4vcic94m-queprmop",
//								"optionName": "",
//								"optionNamePinyin": "",
//								"optionAudioUrl": "",
//								"optionImageUrl": "",
//								"isRightAnswer": 0,
//								"optionMatchId": "",
//								"score": 1
//							}
//						],
//						"questionMatchList": null
//					}
//				]
//			},
//			{
//				"questionProblemId": "4vcic91s-queprm",
//				"questionClassesId": "4v75b5hs-quebankte",
//				"questionClassesName": "TEXT A",
//				"problemDescription": "标题测试A",
//				"problemDescriptionPinyin": "[\"biāo\",\"tí\",\"cè\",\"shì\",\"A\"]",
//				"problemAudioUrl": "",
//				"problemImageUrl": "",
//				"isShowPinyin": 0,
//				"questionSmallProblemList": [
//					{
//						"questionSmallProblemId": "4vcic926-quesmprm",
//						"title": "问题问题问题",
//						"titlePinyin": "",
//						"smallProblemType": 5,
//						"smallProblemDescription": "",
//						"smallProblemDescriptionPinyin": "",
//						"smallProblemAnalysis": "题目解析说明",
//						"score": 1,
//						"audioUrl": "",
//						"imageUrl": "",
//						"questionOptionList": [
//							{
//								"problemOptionId": "4vcic950-queprmop",
//								"optionName": "",
//								"optionNamePinyin": "",
//								"optionAudioUrl": "",
//								"optionImageUrl": "",
//								"isRightAnswer": 0,
//								"optionMatchId": "",
//								"score": 1
//							}
//						],
//						"questionMatchList": null
//					},
//					{
//						"questionSmallProblemId": "4vcic92g-quesmprm",
//						"title": "问题问题问题",
//						"titlePinyin": "[\"wèn\",\"tí\",\"wèn\",\"tí\",\"wèn\",\"tí\"]",
//						"smallProblemType": 4,
//						"smallProblemDescription": "题目文本题目文本题目文本题目文本",
//						"smallProblemDescriptionPinyin": "[\"tí\",\"mù\",\"wén\",\"běn\",\"tí\",\"mù\",\"wén\",\"běn\",\"tí\",\"mù\",\"wén\",\"běn\",\"tí\",\"mù\",\"wén\",\"běn\"]",
//						"smallProblemAnalysis": "题目解析说明题目解析说明题目解析说明",
//						"score": 1,
//						"audioUrl": "",
//						"imageUrl": "",
//						"questionOptionList": [
//							{
//								"problemOptionId": "4vcic95a-queprmop",
//								"optionName": "标准答案标准答案标准答案",
//								"optionNamePinyin": "[\"biāo\",\"zhǔn\",\"dá\",\"àn\",\"biāo\",\"zhǔn\",\"dá\",\"àn\",\"biāo\",\"zhǔn\",\"dá\",\"àn\"]",
//								"optionAudioUrl": "",
//								"optionImageUrl": "",
//								"isRightAnswer": 0,
//								"optionMatchId": "",
//								"score": 1
//							}
//						],
//						"questionMatchList": null
//					},
//					{
//						"questionSmallProblemId": "4vcic92q-quesmprm",
//						"title": "问题问题问题",
//						"titlePinyin": "",
//						"smallProblemType": 1,
//						"smallProblemDescription": "",
//						"smallProblemDescriptionPinyin": "",
//						"smallProblemAnalysis": "题目解析说明",
//						"score": 1,
//						"audioUrl": "",
//						"imageUrl": "",
//						"questionOptionList": [
//							{
//								"problemOptionId": "4vcic95k-queprmop",
//								"optionName": "选项选项选项3",
//								"optionNamePinyin": "[\"xuǎn\",\"xiàng\",\"xuǎn\",\"xiàng\",\"xuǎn\",\"xiàng\",\"3\"]",
//								"optionAudioUrl": "",
//								"optionImageUrl": "",
//								"isRightAnswer": 1,
//								"optionMatchId": "",
//								"score": 1
//							},
//							{
//								"problemOptionId": "4vcic95u-queprmop",
//								"optionName": "选项选项选项2",
//								"optionNamePinyin": "",
//								"optionAudioUrl": "",
//								"optionImageUrl": "",
//								"isRightAnswer": 0,
//								"optionMatchId": "",
//								"score": 1
//							},
//							{
//								"problemOptionId": "4vcic968-queprmop",
//								"optionName": "选项选项选项3",
//								"optionNamePinyin": "",
//								"optionAudioUrl": "",
//								"optionImageUrl": "",
//								"isRightAnswer": 0,
//								"optionMatchId": "",
//								"score": 1
//							}
//						],
//						"questionMatchList": null
//					}
//				]
//			},
//			{
//				"questionProblemId": "4vcic92g-queprm",
//				"questionClassesId": "4v75b5ig-quebankte",
//				"questionClassesName": "TEXT C",
//				"problemDescription": "啊卡就是的卡拉几的是啊卡就是的啦卡就是的",
//				"problemDescriptionPinyin": "[\"ā\",\"kǎ\",\"jiù\",\"shì\",\"de\",\"kǎ\",\"lā\",\"jī\",\"de\",\"shì\",\"ā\",\"kǎ\",\"jiù\",\"shì\",\"de\",\"lā\",\"kǎ\",\"jiù\",\"shì\",\"de\"]",
//				"problemAudioUrl": "",
//				"problemImageUrl": "",
//				"isShowPinyin": 0,
//				"questionSmallProblemList": [
//					{
//						"questionSmallProblemId": "4vcic93o-quesmprm",
//						"title": "大神佳慧大姐卡啊是大姐卡还是的",
//						"titlePinyin": "",
//						"smallProblemType": 1,
//						"smallProblemDescription": "",
//						"smallProblemDescriptionPinyin": "",
//						"smallProblemAnalysis": "啊好开始加大黄金时代卡",
//						"score": 1,
//						"audioUrl": "",
//						"imageUrl": "",
//						"questionOptionList": [
//							{
//								"problemOptionId": "4vcic99c-queprmop",
//								"optionName": "看哈看见的还是看见啊是的",
//								"optionNamePinyin": "[\"kàn\",\"hā\",\"kàn\",\"jiàn\",\"de\",\"huán\",\"shì\",\"kàn\",\"jiàn\",\"ā\",\"shì\",\"de\"]",
//								"optionAudioUrl": "",
//								"optionImageUrl": "",
//								"isRightAnswer": 1,
//								"optionMatchId": "",
//								"score": 1
//							},
//							{
//								"problemOptionId": "4vcic99m-queprmop",
//								"optionName": "啊是客户大姐家卡还得上课",
//								"optionNamePinyin": "",
//								"optionAudioUrl": "",
//								"optionImageUrl": "",
//								"isRightAnswer": 0,
//								"optionMatchId": "",
//								"score": 1
//							},
//							{
//								"problemOptionId": "4vcic9a0-queprmop",
//								"optionName": "看哈看见的还是看见啊是的",
//								"optionNamePinyin": "",
//								"optionAudioUrl": "",
//								"optionImageUrl": "",
//								"isRightAnswer": 0,
//								"optionMatchId": "",
//								"score": 1
//							}
//						],
//						"questionMatchList": null
//					}
//				]
//			}
//		]
//	}`)
//
//	// 32字节的AES密钥
//	key := []byte("Di3cY&_0&7%A1#8-#C839M0-94221X2J")
//
//	// 创建AES加密块
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// 添加填充，确保数据长度为块大小的倍数
//	plaintext = addPadding(plaintext, block.BlockSize())
//
//	// 使用随机生成的IV向量创建一个AES加密流
//	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
//	iv := ciphertext[:aes.BlockSize]
//	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//		panic(err.Error())
//	}
//	stream := cipher.NewCFBEncrypter(block, iv)
//
//	// 加密数据
//	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
//
//	// 将加密后的数据进行Base64编码
//	encrypted := base64.StdEncoding.EncodeToString(ciphertext)
//	fmt.Printf("Encrypted: %s\n", encrypted)
//
//	// 解密
//	decrypted, err := decrypt(encrypted, key)
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Printf("Decrypted: %s\n", decrypted)
//}
//
//// 添加填充，确保数据长度为块大小的倍数
//func addPadding(data []byte, blockSize int) []byte {
//	padding := blockSize - len(data)%blockSize
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(data, padtext...)
//}
//
//// 去除填充
//func removePadding(data []byte) []byte {
//	length := len(data)
//	unpadding := int(data[length-1])
//	return data[:(length - unpadding)]
//}
//
//// 解密
//func decrypt(encrypted string, key []byte) (string, error) {
//	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
//	if err != nil {
//		return "", err
//	}
//
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return "", err
//	}
//
//	if len(ciphertext) < aes.BlockSize {
//		return "", fmt.Errorf("ciphertext too short")
//	}
//
//	iv := ciphertext[:aes.BlockSize]
//	ciphertext = ciphertext[aes.BlockSize:]
//
//	stream := cipher.NewCFBDecrypter(block, iv)
//
//	stream.XORKeyStream(ciphertext, ciphertext)
//
//	// 去除填充
//	plaintext := removePadding(ciphertext)
//
//	return string(plaintext), nil
//}
//
//func Rand32() string {
//	randNumberSeed := []string{
//		"0", "1", "2", "3", "4", "5", "7", "8", "9",
//	}
//	randAlpSeed := []string{
//		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
//		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
//	}
//	randSymSeed := []string{"!", "@", "#", "$", "%", "&", "*", "-", "_", "+", "="}
//	var result string
//	for i := 1; i <= 32; i++ {
//		randNumber := mathRand.Intn(3)
//		var choose string
//		switch randNumber {
//		case 0:
//			randNumberSeedIndex := mathRand.Intn(len(randNumberSeed))
//			choose = randNumberSeed[randNumberSeedIndex]
//		case 1:
//			randAlpSeedIndex := mathRand.Intn(len(randAlpSeed))
//			choose = randAlpSeed[randAlpSeedIndex]
//		case 2:
//			randSymSeedIndex := mathRand.Intn(len(randSymSeed))
//			choose = randSymSeed[randSymSeedIndex]
//		}
//		result += choose
//	}
//	return result
//}
