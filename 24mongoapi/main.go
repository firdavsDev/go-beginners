package main

import (
	"fmt"
	"github.com/mukulmantosh/mongoapi/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Golang with MongoDB")
	fmt.Println("Server is getting started....")
	fmt.Println("Listening on port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
