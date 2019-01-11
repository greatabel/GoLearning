package main

import (
    "fmt"
    "math/rand"
    "time"
)

var era = "AD"

func main() {
    //https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
    rand.Seed(time.Now().UTC().UnixNano())

    year := 2019

    switch month := rand.Intn(12) + 1; month {
    case 2:
        day := rand.Intn(28) + 1
        fmt.Println(era, year, month, day)
    case 4, 6, 9, 11:
        day := rand.Intn(30) + 1
        fmt.Println(era, year, month, day)

    default:
        day := rand.Intn(31) + 1
        fmt.Println(era, year, month, day)
    }


    
}
