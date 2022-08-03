package main

import "fmt"

func main() {
    a := [2]int{0, 0}

    if a == nil {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}