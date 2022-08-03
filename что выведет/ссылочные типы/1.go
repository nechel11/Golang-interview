package main

import (
	"fmt"
)

func main() {
	s := "abc"
	modify(s)
	fmt.Println(s)
}

func modify(a string) {
	a[0] = '3'
}