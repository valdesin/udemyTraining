package main

import "fmt"

func main() {
	if false {
		fmt.Println("First")
	} else if false {
		fmt.Println("Second")
	} else if true {
		fmt.Println("Hello Vlad")
	} else {
		fmt.Println("Third")
	}
}
