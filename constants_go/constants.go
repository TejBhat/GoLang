package main
import "fmt"

const number = 42 //can also be declared outside of main function
// Constants are immutable, meaning they cannot be changed once set.
func main() {
	var name = "Tej"
	name = "sky" // this is allowed because 'name' is now a variable
	fmt.Println(name)
	//output: sky
	fmt.Println(number)
	//output: 42

	const( //this also works
		pi = 3.14
		age = 25
		host="localhost"
	)
	fmt.Println(pi, age, host)
	//output: 3.14 25 localhost
}