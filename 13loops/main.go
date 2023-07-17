package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	//for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}

	//for v := range days {
	//	fmt.Println(v)
	//}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto data
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("Value is ", rogueValue)
		rogueValue++
	}

data:
	fmt.Println("Jumping at Coding")

}
