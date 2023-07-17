package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`              // don't show in the response
	Tags     []string `json:"tags,omitempty"` // if value nil then don't show
}

func main() {
	fmt.Println("Welcome to JSON")
	//fmt.Printf("%s\n", EncodeJson())
	DecodeJson()
}

func EncodeJson() []byte {
	courses := []course{
		{"ReactJS", 299, "google.com", "secure123", []string{"web-dev", "js"}},
		{"MERN", 499, "google.com", "secure12345", []string{"fullstack", "js"}},
		{"Angular", 599, "google.com", "secure1234", nil},
	}
	// package this data as JSON

	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	return finalJson
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
			{
                "course_name": "ReactJS",
                "price": 299,
                "platform": "google.com",
                "tags": ["web-dev", "js"]
			}`)

	var Course course
	checkValidJSON := json.Valid(jsonDataFromWeb)

	if checkValidJSON {
		fmt.Println("Valid JSON")
		json.Unmarshal(jsonDataFromWeb, &Course)
		fmt.Printf("%#v\n", Course)
	} else {
		fmt.Println("Invalid JSON")
	}

	// second case not using struct
	data := make(map[string]interface{})
	json.Unmarshal(jsonDataFromWeb, &data)
	fmt.Printf("%#v\n", data)

	for k, v := range data {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
