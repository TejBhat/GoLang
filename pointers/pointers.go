package main

import "fmt"

func changeNum(num int) {
	num = 10
	fmt.Println("In changeNum", num)
}

func main()  {
	num:= 5
	changeNum(num)
	fmt.Println("After changeNum:", num)
	
}