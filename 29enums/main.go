package main

import "fmt"

// type orderStaus int
type orderStausInString string

// const (
// 	Received orderStaus = iota // automatically sets first value to 0 than second to 1, thrid to 2 ans so on
// 	Confirmed
// 	Prepared
// 	Delivered
// )

const (
	Received  orderStausInString = "received"
	Confirmed orderStausInString = "confirmed"
	Prepared  orderStausInString = "prepared"
	Delivered orderStausInString = "delivered"
)

func changeOrderStatus(status orderStausInString) {
	fmt.Println("Changing order status to", status)
}

func main() {
	// changeOrderStatus(Received)
	changeOrderStatus(Received)
}
