package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Get request in Golang")

	performGetRequest()
}

func performGetRequest() {
	const url = "http://localhost:3000/get"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}
