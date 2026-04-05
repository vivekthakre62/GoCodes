package main

import "fmt"

//this is for integer

// type OrderStatus int

// const (
// 	orderStatus OrderStatus = iota
// 	Confirmed
// 	Pending
// 	Process
// )

//for string
type OrderStatus string

const (
	Received  OrderStatus = "Received"
	Confirmed             = "Confirmed"
	Pending               = "Pending"
	Process               = "Process"
)

func changeOrderStatus(orderStatus OrderStatus) {
	fmt.Println("Order status is:", orderStatus)
}
func main() {
	changeOrderStatus(Confirmed)
}
