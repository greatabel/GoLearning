package main

import (
    "fmt"

)

var era = "AD"

func main() {

    days := 365.2425
    var days1 = 365.2425
    var days2 float64 = 365.2425
    fmt.Println(days, days1, days2)
    // If you initialize a variable with a whole number, 
    // Go won’t know you want floating-point unless you explicitly specify a floating-point type:

    var answer float64 = 42
    fmt.Println(answer)
    
}
