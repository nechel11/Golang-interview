package main

import (
	"fmt"
	"math/rand"
)

func main(){
	a := []int{- 15, 123, 0 ,1233, -124234, 25, 6, 7, 2, 1, 0}
	fmt.Println(QuickSort(a))
}

func QuickSort(a []int) []int {
	len := len(a)
	if len < 2{
		return a
	}
	less := make([]int, 0, len)
	mid := make([]int, 0, len)
	greater := make([]int, 0, len)
	pivot := a[rand.Intn(len)]

	for _,i:=range(a){
		switch{
			case i < pivot:
				less = append(less, i)
			case i == pivot:
				mid = append(mid, i)	
			case i > pivot:
				greater = append(greater, i)
		}
	}
	less, greater = QuickSort(less), QuickSort(greater)
	less = append(less, mid...)
	less = append(less, greater...)
	return less
}
