package main

import (
	"fmt"
	"net/url"
)

const customUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjby27ydrb"

func main() {
	fmt.Println("Welcome to handling URLs")
	fmt.Println(customUrl)

	// parsing
	parse, err := url.Parse(customUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(parse.Scheme)
	fmt.Println(parse.Host)
	fmt.Println(parse.Path)
	fmt.Println(parse.RawQuery)
	fmt.Println(parse.Port())

	params := parse.Query()
	fmt.Println(params)

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=admin",
	}
	fmt.Println(partsOfUrl.String())
}
