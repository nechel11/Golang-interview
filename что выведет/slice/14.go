package main

import (
    "fmt"
)

func main() {
    a := []int{}
    a = append(a, 1)
    a = append(a, 2)
    a = append(a, 3)
    add(a)
    fmt.Println(a)
}

func add(a []int) {
    a = append(a, 4)
}