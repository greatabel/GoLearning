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
    month := rand.Intn(12) + 1
    dayInMonth := 31

    switch month := rand.Intn(12) + 1; month {
    case 2:
        dayInMonth = 28
    case 4, 6, 9, 11:
        dayInMonth = 30
   
    }
    day := rand.Intn(dayInMonth) + 1
    fmt.Println(era, year, month, day)


    
}
