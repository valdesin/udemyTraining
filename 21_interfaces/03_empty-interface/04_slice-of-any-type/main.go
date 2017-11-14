package main

import "fmt"

type animal struct {
	voice string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	tory := dog{animal{"woof"}, true}
	kity := cat{animal{"meow"}, true}
	dogy := dog{animal{"woooof"}, false}

	critters := []interface{}{tory, kity, dogy}
	fmt.Println(critters)
}
