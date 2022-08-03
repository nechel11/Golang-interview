package main

import "fmt"

func main() {
    m := map[[]int]int{
        {1, 2}: 3,
    }

    fmt.Println(m[[]int{1, 2}])
}