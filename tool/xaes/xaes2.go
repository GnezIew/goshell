package xaes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func encrypt(plainText string, key []byte) (string, error) {
	// 创建 AES 分组
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 填充明文以适应 AES 块大小
	blockSize := block.BlockSize()
	plaintext := []byte(plainText)
	plaintext = PKCS5Padding(plaintext, blockSize)

	// 创建加密模式
	iv := make([]byte, blockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}
	stream := cipher.NewCBCEncrypter(block, iv)

	// 加密
	ciphertext := make([]byte, len(plaintext))
	stream.CryptBlocks(ciphertext, plaintext)

	// 将 IV 和密文连接起来
	result := append(iv, ciphertext...)

	// Base64 编码
	return base64.StdEncoding.EncodeToString(result), nil
}

func decrypt(encryptedText string, key []byte) (string, error) {
	// 对 Base64 编码的数据解码
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// 提取 IV
	blockSize := aes.BlockSize
	iv := encryptedData[:blockSize]
	encryptedData = encryptedData[blockSize:]

	// 创建 AES 分组
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建解密模式
	stream := cipher.NewCBCDecrypter(block, iv)

	// 解密
	decrypted := make([]byte, len(encryptedData))
	stream.CryptBlocks(decrypted, encryptedData)

	// 去除填充
	decrypted = PKCS5Unpadding(decrypted)

	return string(decrypted), nil
}

func Xaes2() {
	key := []byte("32-byte-long-key-for-AES-encrypt")
	plainText := "Hello, world!"

	encrypted, err := encrypt(plainText, key)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	fmt.Println("Encrypted:", encrypted)

	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	fmt.Println("Decrypted:", decrypted)

}

// PKCS5Padding 将明文填充为 blockSize 的整数倍
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5Unpadding 去除填充
func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
