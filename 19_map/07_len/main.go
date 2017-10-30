package main

import "fmt"

func main() {
	greeting := map[string]string{
		"Tom":  "Hello",
		"Vlad": "Pryvit",
	}

	greeting["Juan"] = "Hola"

	fmt.Println(len(greeting))
}
