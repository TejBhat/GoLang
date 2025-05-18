package main
import "fmt"
func main() {
	// Declare a variable of type int
	var age int = 25
	fmt.Println(age)

	// Declare a variable of type float64
	var height float64 = 5.9
	fmt.Println(height)

	// Declare a variable of type string
	var name string //= "John Doe" this also works
	fmt.Println(name)

	// Declare a variable of type bool
	var isStudent bool = true
	fmt.Println(isStudent)

	//shorthand syntax
	name = "Tej"
	fmt.Println(name)

	//you dont have to write the data type it automatically detects it
	// this is called type inference
}