package main

import "fmt"

func main() {

	a := Album{}

	fmt.Printf("Album sold : %d\n", a.Sold)
	a = sell(a, 10)

	fmt.Printf("[Main] After sell function Album sold : %d\n", a.Sold)
}

func sell(a Album, qty int) Album {
	a.Sold += qty
	fmt.Printf("[sell] Inside sell function Album sold : %d\n", a.Sold)
	return a
}

type Album struct {
	Name string
	Sold int
}

type Product struct {
	Name     string
	StockQty int
}
