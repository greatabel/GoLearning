package main

import (
    "fmt"

)

func main() {
    // The following two lines are equivalent:
    var count = 110
    count0 := 110
    fmt.Println(count, count0)  
    for count := 10; count > 0; count -- {
        fmt.Println(count)
    }

    fmt.Println(count, count0)
    
}
