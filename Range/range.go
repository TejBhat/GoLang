package main

import "fmt"

//iterating over data structures
func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// for _, num := range nums {
	// 	fmt.Println(num)// prints 1 2 3 4 5
	// }

	// nums := []int{1, 2, 3, 4, 5}
	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println(sum) // prints 15

	// nums := []int{1, 2, 3, 4, 5}
	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// 	}//output: 1 0
	//             // 2 1
	//             // 3 2
	// 	           // 4 3
	//             // 5 4

	m:= map[string]string{"name": "John","city": "New York"}
	for k, v := range m {
		fmt.Println(k, v)
		fmt.Println(k, v)
	}//output: name John
	//         city New York
	}

