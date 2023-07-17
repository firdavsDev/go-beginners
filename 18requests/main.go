package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://example.com"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	fmt.Println(string(dataBytes))
	fmt.Println(response.StatusCode)

}
