package main

import "fmt"

func main(){
	a := 5
	fmt.Println(factorial(a))
}


func factorial(a int) int{
	if a == 1{
		return 1
	} else {
	return a * factorial(a - 1)
	}
}
