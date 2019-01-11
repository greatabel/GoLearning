package main

import "fmt"

func main() {
    var year = 2100
    var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
    if leap {
        fmt.Println("Look before you leap!")
    } else {
        fmt.Println("Keep your feet on the ground.")
    }

    //QC 3.2 answer Which is greater, an “apple” or a “banana”?
    fmt.Println("apple" > "banana")
}