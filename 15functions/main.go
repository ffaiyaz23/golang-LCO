package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")

	// greeter()

	// result := adder(3, 5)
	// fmt.Println("From adder & result is ", result)

	proResult := proAdder(5, 8, 9, 5, 6, 0)
	fmt.Println("From proAdder & proResult is", proResult)
}

func adder(a int, b int) (int, string, string) {
	return a + b, "Added!", "Hello"
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("From greeter")
}
