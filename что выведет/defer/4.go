package main

import (
	"fmt"
)

func main() {
	nums := 1 << 5 // 32

	defer fmt.Println(nums)

	nums = nums >> 1 //16

	fmt.Println("done")
}