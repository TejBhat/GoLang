package main


// func add(a, b int) int {
// 	return a + b
// }

// func main() {
// 	result := add(2, 3)
// 	println(result) // prints 5
// }

// func getLanuages()(string, string, bool) {//when to many return values
// 	return "golang", "python", true
// }

// func main() {
// 	lang1, lang2, lang3 := getLanuages()
// 	println(lang1, lang2, lang3) // prints golang python true
  
// }
func processIt(fn func(a int)int){
	fn(1)
}

func main() {
	fn:=func(a int)int{
		return 2
	}
	processIt(fn)
}//this is incomplete