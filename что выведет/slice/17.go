package main

import "fmt"

func main() {
    a := [2]int{0, 0}

    if a == [2]int{} {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}