package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading..")
	}
	fmt.Println("Thanks for rating: ", input)

}
