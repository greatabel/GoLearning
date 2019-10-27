package main

import (
    "fmt"

    "01Go_In_Action/ch5TypeSystem/i7/counters"
)

func main() {
    // Create a variable of the unexported type and initialize
    // the value to 10.
    counter := counters.New(10)

    // ./listing64.go:15: cannot refer to unexported name
    //                                         counters.alertCounter
    // ./listing64.go:15: undefined: counters.alertCounter

    fmt.Printf("Counter: %d\n", counter)
}