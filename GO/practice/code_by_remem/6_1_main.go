package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	mess := "Some message for writing in the file"

	filepath := "./strange_secret_file.txt"
	file, err := os.Create(filepath)
	ReadErr(err)

	lenn, err := io.WriteString(file, mess)
	fmt.Println("Writed ", lenn)
	ReadErr(err)

	defer file.Close()
	defer ReadFile(filepath)
}

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	ReadErr(err)
	fmt.Println(string(data))
}

func ReadErr(err error) {
	if err != nil {
		panic(err)
	}
}
