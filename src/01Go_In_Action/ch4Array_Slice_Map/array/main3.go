package main

import (
    "fmt"
)

// main is the entry point for the application.
func main() {

    // var array [4][2]int

    array := [4][2]int{{10,11}, {20,21}, {30, 31}, {40, 41}}
    fmt.Printf("There are %d %d \n",array[0][0], array[0][1])
}
