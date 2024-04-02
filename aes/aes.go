package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"strings"
)

// Decrypt 解密
func Decrypt(data string, keyStr string) ([]byte, error) {
	key := []byte(keyStr)
	aesData := strings.Split(data, ":")

	ciphertext0, err := decodeHex(aesData[0])
	if err != nil {
		return nil, err
	}

	ciphertext1, err := decodeHex(aesData[1])
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, ciphertext0)
	crypt := make([]byte, len(ciphertext1))
	blockMode.CryptBlocks(crypt, ciphertext1)

	return pkcs7UnPadding(crypt)
}

func decodeHex(data string) ([]byte, error) {
	return hex.DecodeString(data)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("pkcs7UnPadding error！")
	}

	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

func Encrypt(data []byte, key []byte) (string, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)

	return hex.EncodeToString(key[:blockSize]) + ":" + hex.EncodeToString(crypted), nil
}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
