package main

import "fmt"

const AppType string = "APP"

func main() {
	var username string = "mukul"
	fmt.Println(username)
	fmt.Printf("Variable is of type is %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type is %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type is %T \n", smallVal)

	var smallFloat float32 = 255.2989854788556631
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type is %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type is %T \n", anotherVariable)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)

	// no var style
	numberUser := 30000
	fmt.Println(numberUser)

	fmt.Println(AppType)

}
