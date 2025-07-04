package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int //not public
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

func main() {
	fmt.Println("Methods in Golang")

	// No inheritance in golang, No super, No parent & child
	hitesh := User{"hitesh", "hitesh123@dev.in", false, 30, 15}
	fmt.Println(hitesh)
	fmt.Printf("%+v\n", hitesh)
	hitesh.GetStatus()
}
