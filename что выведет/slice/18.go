package main

import "fmt"

func main() {
    a := [2]int{1, 2}

    if a == [2]int{} {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}