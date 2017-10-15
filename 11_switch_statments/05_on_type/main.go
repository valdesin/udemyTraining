package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func switchOnType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("Type - Int")
	case string:
		fmt.Println("Type - String")
	case contact:
		fmt.Println("Struct Contact")
	default:
		fmt.Println("Unknow")
	}
}

func main() {
	var t = contact{"Hello World", "Vlad"}

	switchOnType(23)
	switchOnType("Vlad")
	switchOnType(t)
	switchOnType(t.greeting)
	switchOnType(t.name)
}
