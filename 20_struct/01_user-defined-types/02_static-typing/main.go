package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 23
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 44
	fmt.Printf("%T %v \n", yourAge, yourAge)

	//s := myAge + yourAge

	//fmt.Println(s)
}
