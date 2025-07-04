package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var myPtr = &myNumber

	fmt.Println("Value of actual pointer is ", myPtr)
	fmt.Println("Value of actual pointer is ", *myPtr)
	fmt.Println("Value of actual pointer is ", &myPtr)
	fmt.Println("Value of actual pointer is ", &myNumber)
}
