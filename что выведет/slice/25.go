package main

import "fmt"

func main() {
    m := map[string]int{}

    if m == map[string]int{} {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}