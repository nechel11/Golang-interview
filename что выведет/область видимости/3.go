package main

import (
    "fmt"
)

func main() {
    a := 5
    switch {
    case a > 1:
        fmt.Println("a > 1")
        fallthrough
    case a > 2:
        fmt.Println("a > 2")
        break
    case a > 3:
        fmt.Println("a > 3")
    default:
        fmt.Println("default")
    }
}