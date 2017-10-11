package main

import "fmt"

func main() {
	a := 10
	b := "Hello"
	c := 5.023
	d := true
	e := 'B'
	f := "Vlad"
	g := `BlablaBlabla`

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
}
