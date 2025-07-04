package main

import "fmt"

func main() {
	fmt.Println("Variables")

	var userName string = "faiz"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T\n", userName)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	age := 15
	fmt.Println(age)

	var sex = "Male"
	fmt.Println(sex)
}
