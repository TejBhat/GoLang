package main

import "fmt"

func changeNum(num int) {
	num = 10
	fmt.Println("In changeNum", num)
}

func maim()  {
	num:= 5
	changeNum(num)
	fmt.Println("After changeNum:", num)
	
}