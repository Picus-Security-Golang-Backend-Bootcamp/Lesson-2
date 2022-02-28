package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := Book{Name: "My First Book"}

	j, err := json.Marshal(a)
	if err != nil {
		fmt.Errorf("JSON dönüştürme işlemi sırasında hata alındı : %s", err.Error())
	}
	fmt.Printf("JSON dönüşümü öncesi : %#v\n", a)
	fmt.Printf("JSON dönüşümü sonrası : %s\n", string(j))

	var request Book

	err = json.Unmarshal(j, &request)

	if err != nil {
		fmt.Errorf("JSON okuma işlemi sırasında hata alındı : %s", err.Error())
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("%#v\n", request)

	request.Name = "My Second Book"

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("%#v\n", request)
}

type Book struct {
	Name    string   `json:"book_name"`
	Reviews []string `json:"reviews,omitempty"`
}
