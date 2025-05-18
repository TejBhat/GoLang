package main
import "fmt"
func main() {
	age := 20
	if age < 18 {
		fmt.Println("You are a minor.")
	}else{
		fmt.Println("You are an adult.")
	}

	//using else if
	if age < 13 {
		fmt.Println("You are a child.")
	}else if age < 18 {
		fmt.Println("You are a teenager.")
	}else if age < 65 {
		fmt.Println("You are an adult.")
	}else{
		fmt.Println("You are a senior citizen.")
	}

	//using logical operators
	var role = "admin"
	var isLoggedIn = true
	if role == "admin" && isLoggedIn {
		fmt.Println("Welcome, Admin!")
	}else if role == "user" && isLoggedIn {
		fmt.Println("Welcome, User!")
	}else{
		fmt.Println("Access Denied!")
	}
	//output: Welcome, Admin!

	//variables declared inside if statement
	if age := 20; age < 18 {
		fmt.Println("You are a minor.")
	}else{
		fmt.Println("You are an adult.")
	}
	//output: You are an adult.

	//Go does not support ternary operator 


}