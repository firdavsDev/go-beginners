package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Files")

	content := "This needs to go in a file - sample Golang"

	file, err := os.Create("./sample.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is :", length)
	defer file.Close()
	ReadFile("./sample.txt")
}

func ReadFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data in file:\n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
