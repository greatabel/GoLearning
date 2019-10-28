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

    array1 := [5]int{1: 10, 2: 20}
    fmt.Printf("There are %d %d %d \n", array1[0], array1[1], array1[2])

    p_array := [5]*int{0: new(int), 1: new(int)}
    *p_array[0] = 100
    *p_array[1] = 200
    fmt.Printf("There are %d %d  \n", *p_array[0], *p_array[1])   
}
