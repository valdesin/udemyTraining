package main

import "fmt"

func main() {
	var s interface{} = "Sydney"
	str, ok := s.(string)
	if ok {
		fmt.Printf("%T \n", str)
	} else {
		fmt.Println("Value is not a string")
	}
}
