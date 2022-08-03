package main

import "fmt"

func main() {
    m := map[string]int{}

    if m == nil {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}