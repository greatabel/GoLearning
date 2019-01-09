package main

import (
    "fmt"
    "math/rand"
)

func main() {
    var degree = 0
    

    for {
        fmt.Println(degree)
        degree++
        var r = rand.Intn(2)
        fmt.Println("r =", r)
        if r == 0 {
            fmt.Println("r =", r, " break")
            break
        }
    }
    
}
