package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer

var mutex sync.Mutex //pointer

func main() {
	//go greeter("Hello")
	//greeter("World")
	website := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
		"https://yahoo.com",
	}

	for _, web := range website {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()
	fmt.Println(signals)

}

//func greeter(s string) {
//	for i := 0; i < 6; i++ {
//		fmt.Println(s)
//	}
//}

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	mutex.Lock()
	signals = append(signals, endpoint)
	mutex.Unlock()
	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}
