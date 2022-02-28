package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generatePassword(max int64) (string, error) {
	b := make([]byte, max)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func main() {
	pass, _ := generatePassword(20)

	fmt.Println(pass)
}
