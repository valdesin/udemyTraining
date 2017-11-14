package main

import "fmt"

func main() {
	var s = "holy"
	str, ok := s.(string)
	if ok {
		fmt.Printf("%q \n", str)
	} else {
		fmt.Println("Value is not a string")
	}
}
