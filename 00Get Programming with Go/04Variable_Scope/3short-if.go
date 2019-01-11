package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    
    if num := rand.Intn(3); num == 0 {
        fmt.Println("Space Adventure")
    } else if num == 1 {
        fmt.Println("SpaceX")
    } else {
        fmt.Println("Virgin Galactic")
    }
    
}
