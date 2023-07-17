package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list of all languages:", languages)
	fmt.Println("JS:", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	for k, v := range languages {
		fmt.Println(k, v)
	}
}
