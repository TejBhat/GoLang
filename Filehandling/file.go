package main

import "os"



func main() {
// 	f,err :=os.Open("test.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fileInfo,err :=f.Stat()
// 	if err != nil {
// 		panic(err)
		
// 	}
// 	fmt.Println("File Name:",fileInfo.Name())// File Name: test.txt
// 	fmt.Println("File Size:",fileInfo.Size())// File Size: 0
// 	fmt.Println("File Mode:",fileInfo.Mode())// File Mode: -rw-r--r--//read ,write
// 	fmt.Println("File IsDir:",fileInfo.IsDir())// File IsDir: false

//Read file

f,err :=os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	println("Read bytes:", n,buf)

 }