package main

import (
	"fmt"
)

func main() {
	var arr [3]string
	arr[0] = "Mas"
	arr[1] = "Maz"
	arr[2] = "Max"
	fmt.Println(arr)

	var numbs = [5]int{1, 2, 5, 6, 7}
	fmt.Println(numbs)
	fmt.Println(len(numbs))
}
