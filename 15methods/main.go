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
	fmt.Println("Is User Active :", mukul.GetStatus())

	mukul.setEmail("abc@gmail.com")
	fmt.Println(mukul.Email)

}

func (u User) GetStatus() bool {
	return u.Status
}

func (u User) setEmail(email string) {
	u.Email = email
	fmt.Println(u.Email)
}
