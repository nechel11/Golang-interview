package main

import (
    "fmt"
)

func main() {
    m := map[string]int{
        "one": 1,
    }

    value := 5

    if value, ok := m["one"]; ok {
        fmt.Println("one exists!")
    }

    fmt.Println(value)
}