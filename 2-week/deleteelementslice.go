package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(deleteItemAndPreserveOrder(arr, 10)) //0-indexed
}

func deleteItem(a []int, i int) []int {
	a[i] = a[len(a)-1]
	a = a[:len(a)-1]

	return a
}

//https://github.com/golang/go/wiki/SliceTricks
func deleteItemAndPreserveOrder(a []int, i int) []int {
	fmt.Println(len(a))
	a = append(a[:i], a[i+1:]...)

	return a
}
