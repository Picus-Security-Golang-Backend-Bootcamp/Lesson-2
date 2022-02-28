package main

import "fmt"

func main() {
	// named type
	var a = struct {
		FirstName string
	}{
		"My First Name",
	}

	var b = struct {
		FirstName string
	}{
		"My First Name",
	}

	c := b

	fmt.Println(a)
	fmt.Println(c)

}
