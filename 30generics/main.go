package main

import (
	"fmt"
)

func printSlice[T any, V string | int](items []T, name V) {
	for _, val := range items {
		fmt.Println(val, name)
	}
}

type Stack[T comparable] struct {
	elements []T
}

func main() {
	printSlice([]int{1, 2, 3}, "Faiz")

	myStack := Stack[string]{
		elements: []string{"Hello", "World"},
	}

	fmt.Println(myStack)
}
