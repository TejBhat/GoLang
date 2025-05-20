package main

import (
	"fmt"
	"time"
)


func task(id int) {
	fmt.Println("  Doing Task", id)
}

func main() {
	for i:=0;i<=10;i++{
		go task(i)
	}

	time.Sleep(2 * time.Second)
	//it runs concurrently
}