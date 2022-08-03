package main

import (
	"fmt"
)

func main() {
	test1 := []int{1, 2, 3, 4, 5}
	test1 = test1[:3]
	test2 := test1[3:]
	fmt.Println(test2[:2])
}
