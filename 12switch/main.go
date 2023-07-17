package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Number 1")
	case 2:
		fmt.Println("Number 2")
	case 3:
		fmt.Println("Number 3")
		fallthrough
	case 4:
		fmt.Println("Number 4")
	case 5:
		fmt.Println("Number 5")
	case 6:
		fmt.Println("Number 6")
	default:
		fmt.Println("Nope!")
	}

}
