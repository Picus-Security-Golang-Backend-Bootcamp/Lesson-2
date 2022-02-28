package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Ahmet", "Mehmet", "Veli", "Büşra", "Ayşe"}

	fmt.Println(s)
	fmt.Println(sort.StringsAreSorted(s)) // false
	sort.Strings(s)
	fmt.Println(sort.StringsAreSorted(s)) // true

	fmt.Println(s)

	v := map[int]struct{} // daha iyi performans, daha kötü memory allocation
	v := map[int]interface{} // daha kötü performans, daha iyi memory allocation
}
