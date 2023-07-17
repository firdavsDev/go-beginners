package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	//fmt.Println(string(content)) alternative solution

	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJSONRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename": "Golang Course",
		"price": 0,
		"platform": "Desktop"
	}
`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}

func PerformPostFORMRequest() {
	const myUrl = "http://localhost:8000/postform"
	// form data

	data := url.Values{}
	data.Add("firstname", "Mukul")
	data.Add("lastname", "Mantosh")
	data.Add("email", "admin@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func main() {
	//PerformGetRequest()
	//PerformPostJSONRequest()
	PerformPostFORMRequest()

}
