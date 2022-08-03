package main

import (
	"fmt"
)

func mod(a []int) {
	a = append(a, 125)
	
	for i := range a {
		a[i] = 5
	}
	
	fmt.Println(a)
}

func main() {
	sl := []int{1, 2, 3, 4}
	mod(sl)
	fmt.Println(sl)
}