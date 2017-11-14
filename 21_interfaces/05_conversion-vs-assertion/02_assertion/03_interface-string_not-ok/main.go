package main

import "fmt"

func main() {
	var x interface{} = 7
	str, ok := x.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("Value is not a string")
	}
}
