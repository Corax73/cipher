package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func encryptFile(key []byte, filepath string, coef int64) error {
	plaintext, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Use CBC mode with random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return err
	}

	cfb := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, aes.BlockSize*coef)
	cfb.CryptBlocks(ciphertext, plaintext)
	ciphertext = append(iv, ciphertext...)

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

	return os.WriteFile(filepath+".enc", []byte(encodedCiphertext), 0644)
}

func decryptFile(key []byte, filepath string, coef int64) error {
	os.ReadFile(filepath)
	ciphertext, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := decodedCiphertext[:aes.BlockSize]
	ciphertext = decodedCiphertext[aes.BlockSize:]

	cfb := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, aes.BlockSize*coef)
	cfb.CryptBlocks(plaintext, ciphertext)

	return os.WriteFile(filepath+".enc", plaintext, 0644)
}

func main() {
	// Replace with your desired password
	password := "1234567890123456"
	key := []byte(password)

	// Specify the file to encrypt/decrypt
	filepath := "./data/your_file.txt"

	// Choose between encryption or decryption
	action := "encrypt" // or "decrypt"

	file, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("File not found:", err)
	}
	coef := file.Size()
	if file.Size() > aes.BlockSize {
		coef = (file.Size() % aes.BlockSize)
	}
	coef += 100

	if action == "encrypt" {
		err := encryptFile(key, filepath, coef)
		if err != nil {
			fmt.Println("Error encrypting file:", err)
			return
		}
		fmt.Println("File encrypted successfully!")
	} else if action == "decrypt" {
		err := decryptFile(key, filepath+".enc", coef)
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
		fmt.Println("File decrypted successfully!")
	} else {
		fmt.Println("Invalid action. Please choose 'encrypt' or 'decrypt'.")
	}
}
