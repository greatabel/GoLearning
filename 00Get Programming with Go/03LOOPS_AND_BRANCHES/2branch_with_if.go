package main

import "fmt"

func main() {
    var command = "go east"
    if command == "go east" {
        fmt.Println("You head further up the mountain.")
    } else {
        fmt.Println("Did't quite get that.")
    }

    //QC 3.2 answer Which is greater, an “apple” or a “banana”?
    fmt.Println("apple" > "banana")
}