package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Get request in Golang")

	performGetRequest()
	performPostRequest()
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

func performPostRequest() {
	const myUrl = "http://localhost:3000/post"

	//fake json Payload
	requestBody := strings.NewReader(`
		{
			"course": "golang",
			"price": 599,
			"platform": "learncode.online.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func performPostFormRequest() {
	myUrl := "http://localhost:3000/postform"

	//making formdata
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "singh")
	data.Add("firstname", "abc@gmail.com")

	response, _ := http.PostForm(myUrl, data)
	content, _ := io.ReadAll(response.Body)

	defer response.Body.Close()

	fmt.Println(string(content))
}
