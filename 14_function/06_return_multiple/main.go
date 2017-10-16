package main

import "fmt"

func main() {
	fmt.Println(green("Vllad ", "Fil "))
}

func green(name, lastName string) (string, string) {
	return fmt.Sprint(name, lastName), fmt.Sprint(lastName, name)
}
