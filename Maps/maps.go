package main

import (
	"fmt"
	"maps"
)

//maps-> hash,object,dictionary
func main() {
	//creating map
	m := make(map[string]string)
	//setting element
	m["name"] = "Tej"
	m["area"] = "India"
	//getting element
	fmt.Println(m["name"], m["area"])
	//output: Tej India

//imp->if key does not exist in the map than it returns the zero value of the type
	fmt.Println(m["age"]) //output: ""

	ageMap := make(map[string]int)
	ageMap["age"] = 25
	ageMap["height"] = 5
	fmt.Println(ageMap["age"]) //output: 25
	fmt.Println(len(ageMap)) //output: 1

	//deleting element
	delete(ageMap, "age")
	fmt.Println(ageMap)//output: map[height:5]
	
	//clear
	clear(ageMap)
	fmt.Println(ageMap)//output: map[]

	//creating map without make keyword
	// m2 := map[string]int{"name": 1, "age": 2}
	// fmt.Println(m2) //output: map[age:2 name:1]

	// _, ok := m2["name"]
	// if ok {
	// 	fmt.Println("Key exists")
	// } else {
	// 	fmt.Println("Key does not exist")
	// }//output: Key exists

	// v, ok := m2["name"]
	// if ok {
	// 	fmt.Println("Key exists with value:", v)
	// } else {
	// 	fmt.Println("Key does not exist")
	// }//output: Key exists with value: 1

	m1 := map[string]int{"name": 1, "age": 2}
	m2 := map[string]int{"name": 1, "age": 2}
	fmt.Println(maps.Equal(m1, m2)) //output: true

}