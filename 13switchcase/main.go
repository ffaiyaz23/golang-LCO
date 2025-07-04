package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case in Golang")

	// rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
		fallthrough
	case 4:
		fmt.Println("Number is 4")
		fallthrough
	case 5:
		fmt.Println("Number is 5")
	case 6:
		fmt.Println("Number is 6")

	}
}
