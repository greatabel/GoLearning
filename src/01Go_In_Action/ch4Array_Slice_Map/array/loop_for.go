package main

import (
    "fmt"
)

// main is the entry point for the application.
func main() {

    slice := []int{10, 20, 30, 40, 50}
    for index := 2; index < len(slice); index++ {
        fmt.Printf("Index: %d Value: %d\n", index, slice[index])
    }





}
