package plugins

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var (
	aesKey []byte
)

func NewAesKey(key string) {
	t := len(key)
	if t < 16 {
		for i := t; i < 16; i++ {
			key += "."
		}
	} else {
		key = key[0:16]
	}
	aesKey = []byte(key)
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func aesEncrypt(value string) string {
	origData := []byte(value)
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	blockSize := block.BlockSize()
	origData = pKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, aesKey[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.URLEncoding.EncodeToString(crypted)
}
func aesDecrypt(value string) (string, error) {
	crypted, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, aesKey[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	return string(pKCS7UnPadding(origData)), nil
}

/*******************加密**********************/
func Aes_Encode(value string) string {
	return aesEncrypt(value)
}

/*******************解密**********************/
func Aes_Decode(value string) (string, error) {
	return aesDecrypt(value)
}
