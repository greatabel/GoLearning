// Quick check 2.6 The distance between Earth and Mars varies 
// from nearby to opposite sides of the sun. 
// Write a program that generates a random distance from 56,000,000 to 401,000,000 km.

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
    rand.Seed(time.Now().UTC().UnixNano())
    // a random distance to Mars (km)
    var distance = rand.Intn(345000001) + 56000000
    fmt.Println(distance)
}