package main

import "fmt"

func main(){
	a := []int{6,7,4,8,0,121,4,5,}
	fmt.Println("after sort", a)
	select_sort(a)
	fmt.Println("before sort", a)
	
}


func select_sort(a []int){
	len := len(a)

	for i:=0; i < len; i++{
		minIndx := i
		for j := i; j < len; j++{
			if a[j] < a[minIndx]{
				minIndx = j
			}
		}
		a[i], a[minIndx] = a[minIndx], a[i]
	}
}