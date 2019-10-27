// Sample program to show how to show you how to briefly work with io.
package main

import (
    "fmt"


)

// main is the entry point for the application.
func main() {

    var array1 [5]string

    array2 := [5]string{"red", "blue", "green", "yellow", "pink"}

    array1 = array2
    fmt.Printf("There are %s %s %s \n", array1[0], array1[1], array1[2])

    // 指针数组赋值给另一个
    var array3 [3]*string

    array4 := [3]*string{new(string), new(string), new(string)}
    *array4[0] = "Red"
    *array4[1] = "Green"
    *array4[2] = "Blue"
    array3 = array4
    fmt.Printf("There are %s %s %s \n", *array3[0], *array3[1], *array3[2])
}
