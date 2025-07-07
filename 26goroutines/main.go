package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup // these are pointers
var mutex sync.Mutex         // pointer

var signals = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://www.linkedin.com",
		"https://www.golangjobs.tech",
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		waitGroup.Add(1)
		go getStatusCode(web)
	}

	waitGroup.Wait()

	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	result, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)

	mutex.Lock()
	signals = append(signals, endpoint)
	mutex.Unlock()

	defer waitGroup.Done()
}
