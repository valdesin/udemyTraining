package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.ParseInt("-234", 10, 64)
	b, _ := strconv.ParseUint("456", 10, 64)
	c, _ := strconv.ParseFloat("23.35464", 64)
	d, _ := strconv.ParseBool("true")

	fmt.Println(a, b, c, d)
	fmt.Printf("%T %T %T %T \n", a, b, c, d)

	e := strconv.FormatInt(-645, 10)
	f := strconv.FormatUint(653, 10)
	g := strconv.FormatFloat(152.36635, 'E', -1, 64)
	k := strconv.FormatBool(true)

	fmt.Println(e, f, g, k)
	fmt.Printf("%T %T %T %T \n", e, f, g, k)
}
