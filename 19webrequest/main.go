package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://pkg.go.dev/"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T", response)
	// fmt.Println(response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))
}
