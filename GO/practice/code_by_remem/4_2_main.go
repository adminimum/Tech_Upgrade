package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var result string

	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "Monday"
	case 2:
		result = "Thuesday"
		fallthrough
	default:
		result = "Not Monday"
	}
	fmt.Println(result)
}
