package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2}
	y := []int{3, 4}
	ref := x
	x = y
	fmt.Println(x, y, ref)
}