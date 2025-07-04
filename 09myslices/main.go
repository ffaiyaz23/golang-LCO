package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitLess = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitLess is %T\n", fruitLess)

	fruitLess = append(fruitLess, "Mango", "Banana")
	fmt.Println(fruitLess)

	// fruitLess = append(fruitLess[1:])
	fmt.Println(fruitLess)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867

	highScore = append(highScore, 555, 666, 321)
	sort.Ints(highScore)

	fmt.Println(highScore)
}
