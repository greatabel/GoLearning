package main

import (
    "fmt"
)

// main is the entry point for the application.
func main() {

    // dict := make(map[string]int)

    dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

    dict["FakeColor"] = "#fake111"
    // for index := 2; index < len(slice); index++ {
    //     fmt.Printf("Index: %d Value: %d\n", index, slice[index])
    // }
    fmt.Printf("%s haha\n", dict["FakeColor"])





}
