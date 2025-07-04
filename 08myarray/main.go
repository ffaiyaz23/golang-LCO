package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Mango"
	fruitList[3] = "Lichi"

	fmt.Println("FruitList is : ", fruitList)

	var veggieList = [3]string{"Potato", "tomato", "brinjal"}
	fmt.Println("VeggieList is : ", veggieList)
}
