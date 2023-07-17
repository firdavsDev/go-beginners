package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions!")
	greeter()
	display()

	result := adder(4, 3)
	fmt.Println("Result is: ", result)

	finalResult := proAdder(1, 2, 3, 4)
	fmt.Println("Final Result is: ", finalResult)

}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func adder(i int, i2 int) int {
	return i + i2
}

func greeter() {
	fmt.Println("Testing")
}

func display() {
	fmt.Println("Inside Display")
}
