package main

import (
	"fmt"
)

func main() {
	dd := Dog{"Woof", 12, "Poodle"}

	dd.Speak()
	dd.ThreeSpeak()

}

type Dog struct {
	Bark   string
	Weight int
	Breed  string
}

func (d Dog) Speak() {
	fmt.Println(d.Bark)
}

func (d Dog) ThreeSpeak() {
	d.Bark = fmt.Sprintf("%v %v %v", d.Bark, d.Bark, d.Bark)
	fmt.Println(d.Bark)
}
