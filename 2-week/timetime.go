package main

import (
	"fmt"
	"time"
)

func main() {
	// Go'da date ve time aynı şeyi ifade ediyor.
	// Unix Epoch Time 1 Ocak 1970 00:00'dan şu ana kadar geçen saniyeyi döner
	now := time.Now()
	fmt.Println(now)
	location, _ := time.LoadLocation("Europe/London")

	fmt.Println(now.In(location))
}
