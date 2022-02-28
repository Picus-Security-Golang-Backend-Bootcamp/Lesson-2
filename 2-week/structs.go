package main

import "fmt"

func main() {

	// https://medium.com/a-journey-with-go/go-introduction-to-the-escape-analysis-f7610174e890

	// https://itnext.io/structure-size-optimization-in-golang-alignment-padding-more-effective-memory-layout-linters-fffdcba27c61

	//maligned
	//a := Person{"asdsaf", "asdaf", 12, 23}

	c := new(Person)

	b := Person{
		FirsrtName: "asadfs",
		LastName:   "ewfsdf",
		Age:        23,
		Salary:     30,
	}

	//fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

type Person struct {
	Age        int
	Salary     float64
	FirsrtName string
	LastName   string
}
