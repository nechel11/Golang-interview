package main

import "fmt"

func main() {
    a := []int{}

    if a == nil {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}