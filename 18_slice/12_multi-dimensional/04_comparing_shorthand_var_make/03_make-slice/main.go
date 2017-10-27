package main

import "fmt"

func main() {
	student := make([]string, 14)
	students := make([][]string, 14)

	student[0] = "Vlad"
	//student = append(student, "Vlad")
	fmt.Println(student)
	fmt.Println(students)
}
