package main

import (
	"fmt"
)

func main(){
	a := []int{1,2,3,4,5,6,7,8,9}
	b := 4
	fmt.Println("Searchin int is", b)
	fmt.Println("Index num is",binary_search(a,b))
}

func binary_search(a []int, b int) int{
	low := 0
	high := len(a) - 1

	
	for low <= high{
		mid := (high + low) / 2
		guess := a[mid]
		if guess == b{
			return mid
		} else if guess < b {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0
}