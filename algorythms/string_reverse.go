package main

import "fmt"

func main(){
	s := "ABCD"
	s1 := "QWERTyuiop"
	fmt.Println(reverse(s), reverse(s1))
}

func reverse(s string) string{
	var res string

	for _, v := range(s){
		res = string(v) + res
	}
	return res
}