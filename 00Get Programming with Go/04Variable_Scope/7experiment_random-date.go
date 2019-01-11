package main

import (
    "fmt"
    "math/rand"
    "time"
)

var era = "AD"

func main() {
    var count = 10 
    for count > 0 {
        count--

        rand.Seed(time.Now().UTC().UnixNano())

        year := rand.Intn(2019) + 1
        month := rand.Intn(12) + 1
        dayInMonth := 31

        switch month := rand.Intn(12) + 1; month {
        case 2:
            var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
            if leap {
                dayInMonth = 29
            } else {
                dayInMonth = 28
            }

        case 4, 6, 9, 11:
            dayInMonth = 30
       
        }
        day := rand.Intn(dayInMonth) + 1
        fmt.Println(era, year, month, day)
    }



    
}
