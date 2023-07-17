package main

import "fmt"

func main() {

	fmt.Println("Welcome to If-Else")

	loginCount := 25
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 15 && loginCount <= 24 {
		result = "Super User"
	} else {
		result = ":("
	}

	fmt.Println(result)

}
