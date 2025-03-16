package main

import (
	"cipher/core"
	"fmt"
)

func main() {
	password := "+W^apDG_nxq0PKJa"
	key := []byte(password)

	filepath := "./data/your_file.txt"

	action := "encrypt" // or "decrypt"

	if action == "encrypt" {
		err := core.EncryptFile(key, filepath)
		if err != nil {
			fmt.Println("Error encrypting file:", err)
			return
		}
		fmt.Println("File encrypted successfully!")
	} else if action == "decrypt" {
		err := core.DecryptFile(key, filepath+".enc")
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
		fmt.Println("File decrypted successfully!")
	} else {
		fmt.Println("Invalid action. Please choose 'encrypt' or 'decrypt'.")
	}
}
