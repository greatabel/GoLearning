package main

import (
    "fmt"
)

// main is the entry point for the application.
func main() {

    slice := []int{10, 20, 30, 40}
    for index, value := range slice {
        fmt.Printf("Index: %d Value: %d\n", index, value)
    }



}
