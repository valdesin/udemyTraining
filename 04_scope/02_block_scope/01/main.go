package main

import "fmt"

func main() {
	x := 100
	fmt.Println(x)
	fool()
}

func fool() {
	// no access to x
	// this does not compile
	//fmt.Println(x)
}
