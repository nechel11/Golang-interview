package main

import (
	"fmt"
)

type book struct {
	author string
	title  string
}

func (a book) String() string {
	return a.author + a.title
}

type Chislo int

func (a Chislo) String() string {
	return "hello world"
}

func write(s fmt.Stringer) {
	fmt.Println(s)
}

func main() {
	a := book{"zafar", "kahsanov"}
	b := Chislo(25)

	write(a)
	write(b)
}
