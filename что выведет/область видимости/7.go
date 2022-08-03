package main

import (
    "fmt"
)

func main() {
    m := map[string]int{
        "zero": 0,
    }

    value := 5

    if value, ok := m["one"]; ok {
        fmt.Println(value)
    }

    fmt.Println(value)
}