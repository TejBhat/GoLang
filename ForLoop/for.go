package main
import "fmt"
// The for loop is the only loop in Go.
func main() {
	//this does same as while loop- there is no while loop in go
	i:=1
	for i<=5 {
		fmt.Println(i)
		i++
	}
	//infinite loop
	// for {
	// 	Println("Hello, World!")	
	// }

	//classic for loop
	for i:=1; i<=5; i++ {
		fmt.Println(i)
	}
	//break and continue
	for i:=1; i<=5; i++ {
		if i==3 {
			continue //skip this iteration
		}
		if i==5 {
			break //exit the loop
		}
		fmt.Println(i)
	}
	//output: 1 2 4

	//range loop
	//range loop is used to iterate over arrays, slices, maps, strings, and channels
	for i := range 10 {
		fmt.Println(i)
	}
	//output: 0 1 2 3 4 5 6 7 8 9
	
}