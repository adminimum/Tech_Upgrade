package main

import (
	"fmt"
)

func main() {
	anInt := 22
	var pp = &anInt
	fmt.Println("An int", *pp)

	val := 12
	ppp := &val
	fmt.Println("val =", *ppp)

	*ppp = *ppp / 3

	fmt.Println("ppp/3 =", *ppp)
	fmt.Println("val =", val)
}
