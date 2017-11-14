package main

import "fmt"

func main() {
	var x interface{} = 7
	fmt.Println(x.(int) + 6)
}
