package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(i interface{}) {
	fmt.Println(i)
}

func main() {
	d := dog{animal{"woof"}, true}
	c := cat{animal{"meow"}, true}

	specs(d)
	specs(c)
}
