package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		defer func(i *int) {
			fmt.Printf("%v ", *i)
		}(&i)
	}

}