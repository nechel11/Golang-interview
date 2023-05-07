package main

import "fmt"

func merge(arrs ...[]int) []int {
	res := make([]int, 0, len(arrs))
	fmt.Println()

	for _, v := range arrs {
		res = append(res, v...)
	}
	return res
}

func main() {
	fmt.Println(merge([]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8, 9, 10}))
}
