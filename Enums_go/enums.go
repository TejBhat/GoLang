package main

import "fmt"

//enumerated types are a way to define a set of named values that are related to each other.
// They are often used to represent a finite set of options or states.

//type OrderStatus int
// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delievered
// )//iota only for integer types

type OrderStatus string
const (
	Received OrderStatus = "Received"
	Confirmed            = "Confirmed"
	Prepared             = "Prepared"
	Delievered           = "Delievered"
)


func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing status changed to:", status)
}
func main() {
	changeOrderStatus(Prepared)// prints Order status changed to: 2
}
