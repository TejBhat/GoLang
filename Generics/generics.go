package main

import "fmt"



func printSlice[T any](items []T) { //only for allowinf int and string use this // func printSlice[T int | string](items []T) 
	//can also use [T interface{}] to allow any type
	//can also use [T comparable] to allow any type that is comparable
	for _, item := range items {
		fmt.Println(item)
	}
}



func main() {
	nums := []int{1, 2, 3, 4, 5}
	printSlice(nums)// prints 1 2 3 4 5
}