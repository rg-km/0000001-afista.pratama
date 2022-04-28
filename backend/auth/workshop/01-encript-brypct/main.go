package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hashedStr, err := encryptToBcrypt("hello")
	if err != nil {
		panic(err)
	}

	//bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte("hello"))

	fmt.Printf("Hashed String = %s", hashedStr)
}

// encrypt string kedalam bcrypt
func encryptToBcrypt(str string) (string, error) {
	// Task: Hashing the password with the default cost of 10

	encrypted, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encrypted), nil // TODO: replace this
}
