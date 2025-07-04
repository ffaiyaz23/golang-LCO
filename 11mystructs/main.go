package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in Golang")

	// No inheritance in golang, No super, No parent & child
	hitesh := User{"hitesh", "hitesh123@dev.in", false, 30}
	fmt.Println(hitesh)
	fmt.Printf("%+v", hitesh)
}
