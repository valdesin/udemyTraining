package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = "12"
	var y = 12
	z, _ := strconv.Atoi(x)
	fmt.Println(z + y)
	fmt.Println(z)
}
