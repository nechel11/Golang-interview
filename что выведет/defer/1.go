package main

import (
	"fmt"
)

func main() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}