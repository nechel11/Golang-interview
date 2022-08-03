package main

import "fmt"

func main() {
    a := []int{}

    if a == []int{} {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}