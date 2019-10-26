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

  
}
