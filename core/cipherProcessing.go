package core

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"os"
)

func EncryptFile(key []byte, filepath string) (string, error) {
	plaintext, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	end := len(plaintext) % aes.BlockSize
	if end != 0 {
		countSteps := aes.BlockSize - end
		for range countSteps {
			plaintext = append(plaintext, 0)
		}
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.CryptBlocks(ciphertext, plaintext)
	ciphertext = append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(ciphertext), err
}

func DecryptFile(key []byte, filepath string) ([]byte, error) {
	emptyResp := make([]byte, 0)
	os.ReadFile(filepath)
	ciphertext, err := os.ReadFile(filepath)
	if err != nil {
		return emptyResp, err
	}
	decodedCiphertext, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return emptyResp, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return emptyResp, err
	}
	iv := decodedCiphertext[:aes.BlockSize]
	ciphertext = decodedCiphertext[aes.BlockSize:]
	cfb := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(decodedCiphertext))
	cfb.CryptBlocks(plaintext, ciphertext)
	plaintext = deletionFillChars(plaintext)
	return plaintext, err
}

func deletionFillChars(src []byte) []byte {
	length := len(src)
	var unPadding int
	for _, val := range src {
		if rune(val) == rune(0) {
			unPadding++
		}
	}
	return src[:(length - unPadding)]
}
