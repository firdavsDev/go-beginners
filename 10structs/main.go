package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to Structs")
	mukul := User{Name: "Mukul", Email: "mukul@gmail.com", Status: true, Age: 25}
	fmt.Println(mukul)
	fmt.Printf("%+v\n", mukul)

}
