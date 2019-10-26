// Sample program to show how to show you how to briefly work with io.
package main

import (
    "fmt"


)

// main is the entry point for the application.
func main() {
    // var array [5]int

    // array := [5]int{10, 20, 30, 40, 50}

    // 容量初始值决定
    array := [...]int{100, 200, 300}
    fmt.Printf("There are %d %d %d \n", array[0], array[1], array[2])
}
