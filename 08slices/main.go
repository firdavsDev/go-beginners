package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fruitList = append(fruitList, "mango", "Banana")
	fmt.Printf("%T\n", fruitList)
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 2321
	highScores[1] = 2322
	highScores[2] = 2323
	highScores[3] = 2324

	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// remove value from slice based on index

	var courses = []string{"react", "vue", "angular", "swift", "js", "python"}

	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
