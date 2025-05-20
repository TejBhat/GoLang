package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time// time.Time is a struct in the time package
	// time.Time is a struct that represents a point in time
	// it has methods like Add, Sub, etc. to manipulate time
	// it also has methods like Year, Month, Day, etc. to get the date
	//nanosecond precision
}
 func main() {
	myOrder := order{
		id: "123",
		amount : 100.0,
		status: "pending",
 }

   myOrder.createdAt = time.Now()// time.Now() returns the current time
   // time.Now() is a function that returns the current time
   //prints-> 2023-10-01 12:00:00 +0000 UTC
   fmt.Println(myOrder.status)//output: 2023-10-01 12:00:00 +0000 UTC
 fmt.Println("order struct",myOrder)//output: order struct {123 100 pending 0001-01-01 00:00:00 +0000 UTC}
}