package main

import "fmt"

func main() {
	student := make([]string, 32)
	students := make([][]string, 32)

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
