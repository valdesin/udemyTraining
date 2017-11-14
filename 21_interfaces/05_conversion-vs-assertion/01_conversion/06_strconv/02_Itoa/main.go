package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 12
	var s = "I have many " + strconv.Itoa(x)
	fmt.Println(s)
}
