package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	var fruits []string

	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Orange")
	fruits = append(fruits, "Pineapple")

	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var vegList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println(vegList)

}
