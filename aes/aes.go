package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"strings"
)

// AesDecrypt 解密
func AesDecrypt(data string, keyStr string) ([]byte, error) {
	key := []byte(keyStr)
	aesData := strings.Split(data, ":")

	ciphertext_0, decodeString_0_err := hex.DecodeString(aesData[0])
	if decodeString_0_err != nil {
		return nil, decodeString_0_err
	}

	ciphertext_1, decodeString_1_err := hex.DecodeString(aesData[1])
	if decodeString_1_err != nil {
		return nil, decodeString_1_err
	}

	block, newCipherErr := aes.NewCipher(key)

	if newCipherErr != nil {
		return nil, newCipherErr
	}

	blockMode := cipher.NewCBCDecrypter(block, ciphertext_0)

	crypted := make([]byte, len(ciphertext_1))

	blockMode.CryptBlocks(crypted, ciphertext_1)

	crypted, _ = pkcs7UnPadding(crypted)

	return crypted, nil
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("pkcs7UnPadding error！")
	}

	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}
