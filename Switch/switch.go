package main

import "time"

func main() {
	// Switch statement
	var day int = 3
	switch day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid day")
	}
	//output: Wednesday

	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend!")
	default:
		println("It's a weekday.")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("I am an int")
		case string:
			println("I am a string")
		case bool:
			println("I am a bool")
		default:
			println("I don't know what I am")
		}
	}
	whoAmI(42)        // I am an int
	whoAmI("Hello")   // I am a string
}