package main

import (
    "fmt"
)

func main() {
    a := 2

    switch a {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
        fallthrough
    case 3:
        fmt.Println("three")
        fallthrough
    default:
        fmt.Println("default")
    }
}