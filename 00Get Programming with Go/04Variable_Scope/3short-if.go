package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    //https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
    rand.Seed(time.Now().UTC().UnixNano())

    if num := rand.Intn(3); num == 0 {
        fmt.Println("Space Adventure")
    } else if num == 1 {
        fmt.Println("SpaceX")
    } else {
        fmt.Println("Virgin Galactic")
    }
    
}
