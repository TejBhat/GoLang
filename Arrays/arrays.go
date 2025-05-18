package main

import "fmt"

func main() {
	// var nums [4]int
	// nums[0] = 1
	// fmt.Println(len(nums)) 

	// var vals [4]bool
	// vals[2]= 
	// fmt.Println(vals[2]) // prints false

	//to declare in single line
	nums:= [4]int{1, 2, 3, 4}
	fmt.Println(nums) // prints [1 2 3 4]

	//2D array
	// 2D array
	nums2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums2D) // prints [[1 2 3] [4 5 6]]
}

// fixed size, that is predicable
// memory optimization
// const time access