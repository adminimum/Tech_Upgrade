package main

import (
	"fmt"
	"sort"
)

func main() {

	states := make(map[string]string)
	states["MA"] = "Mander Aoure"
	states["KC"] = "Kepo Cale"
	states["LE"] = "Leviro Erdono"
	states["SS"] = "Sarto Samin"

	fmt.Println(states)

	delete(states, "SS")

	fmt.Println(states)

	states["PO"] = "Pontae Omeno"

	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	ind := 0

	for k := range states {
		keys[ind] = k
		ind++
	}
	sort.Strings(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
