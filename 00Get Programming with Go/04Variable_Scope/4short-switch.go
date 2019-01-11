package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    //https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
    rand.Seed(time.Now().UTC().UnixNano())

    switch num := rand.Intn(10); num {
    case 0:
     fmt.Println("Space Adventures")
    case 1:
     fmt.Println("SpaceX")
    case 2:
     fmt.Println("Virgin Galactic")
    default:
     fmt.Println("Random spaceline #", num)
    }


    
}
