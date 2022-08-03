package main

import "fmt"

func main() {
    m := map[[2]int]int{
        {1, 2}: 3,
    }

    fmt.Println(m[[2]int{1, 2}])
}