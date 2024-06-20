package main

import (
	"fmt"
	"sort"
)

func main() {
	// var arr [3]string
	// arr[0] = "Mas"
	// arr[1] = "Maz"
	// arr[2] = "Max"
	// fmt.Println(arr)

	var arr = []string{"Mas", "Maz", "Max"}
	fmt.Println(arr)

	arr = append(arr, "Mar")
	fmt.Println(arr)

	arr = append(arr[1:len(arr)])
	fmt.Println(arr)

	arr = append(arr[:len(arr)-1])
	fmt.Println(arr)

	numbs := make([]int, 5) // 	numbs := make([]int,5,5) restrict append values
	numbs[0] = 234
	numbs[1] = 27123
	numbs[2] = 21
	numbs[3] = 54
	numbs[4] = 78
	fmt.Println(numbs)

	numbs = append(numbs, 989)
	sort.Ints(numbs)
	fmt.Println(numbs)

}
