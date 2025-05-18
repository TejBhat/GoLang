package main

import (
	"fmt"
	"slices"
)

//slice is a dynamic array
//most used construct in go

func main() {
	//uninitialized slice is nil
	//var nums []int
	//fmt.Println(nums==nil) // prints []

	var nums2 =make([]int, 5) // creates a slice of length 5
	//fmt.Println(nums2) // prints [0 0 0 0 0]

	//cap-capacity->maximum number of elements that can be stored in the slice
	fmt.Println(cap(nums2)) // prints 5

	nums2 = append(nums2, 1) // adds 1 to the slice
	fmt.Println(nums2) // prints [0 0 0 0 0 1]

	//copy function
	var nums3 =make([]int,0, 5) // creates a slice of length 0 and capacity 5
	nums3 = append(nums3, 2)
	var nums4 =make([]int, len(nums3)) // creates a slice of length 5

	fmt.Println(nums3,nums4) // prints [2] [0 0 0 0 0]
	copy(nums4, nums3)
	//output: [2] [2 0 0 0 0]

	//slice operator
	var num = []int{1,2,3}
	fmt.Println(num[0:2]) // prints [1 2]
	fmt.Println(num[1:]) // prints [2 3]
	fmt.Println(num[:2]) // prints [1 2]

	//slice package - inbuilt
	var nums5 = []int{1,2,3,4,5}
	var nums6 = []int{6,7,8,9,10}
	fmt.Println(slices.Equal(nums5, nums6)) // prints false
	fmt.Println(slices.Contains(nums5, 3)) // prints true

	//also this can be done
	var nums7 = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums7) // prints [[1 2 3] [4 5 6]]
}