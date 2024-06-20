package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("Time now", n)

	t := time.Date(20011, time.November, 11, 0, 0, 0, 0, time.UTC)
	fmt.Println("Spec time", t)
	fmt.Println("Format ANSIC", t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Fri Nov 12 00:01:00 2021")
	fmt.Printf("Type of str Date%T\n", parsedTime)
}
