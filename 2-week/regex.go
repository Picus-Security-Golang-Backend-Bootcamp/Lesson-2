package main

import (
	"fmt"
	"regexp"
)

func main() {
	rex, _ := regexp.Compile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")

	// 2. yöntem

	//isMatch, _ := regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", "denemedeneme.dsfsdf.sf.com")

	email := "denemedeneme.dsfsdf.sf.com"

	isMatch := rex.MatchString(email)

	if !isMatch {
		fmt.Printf("Girilen e-posta adresi geçerli bir formata sahip değil: %s\n", email)
		return
	}

	fmt.Println("E-posta formatı doğru")
}
