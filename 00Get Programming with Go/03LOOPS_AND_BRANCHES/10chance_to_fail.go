package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    var count = 15
    

    for count > 0 {
        fmt.Println("count= ",count)
        time.Sleep(time.Second)
        var r = rand.Intn(100)
        fmt.Println("r =", r)
        if r == 0 {
            fmt.Println("r =", r, " break")
            break
        }
        count--
    }
    if count == 0 {
        fmt.Println("Liftoff!")
    } else {
        fmt.Println("Launch failed")
    }
    
}
