package core

import "math/rand"

func PasswordGenerator(passwordLength int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specialChar := "!@#$%^&*()_-+={}[/?]"
	password := ""

	for range passwordLength {
		randNum := rand.Intn(4)
		switch randNum {
		case 0:
			randCharNum := rand.Intn(len(lowerCase))
			password += string(lowerCase[randCharNum])
		case 1:
			randCharNum := rand.Intn(len(upperCase))
			password += string(upperCase[randCharNum])
		case 2:
			randCharNum := rand.Intn(len(numbers))
			password += string(numbers[randCharNum])
		case 3:
			randCharNum := rand.Intn(len(specialChar))
			password += string(specialChar[randCharNum])
		}
	}

	return password
}
