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

    value, exists := dict["Blue"]
    if exists {
        fmt.Printf(value)
    }

    valueR := dict["Red"]
    if valueR != "" {
        fmt.Println(valueR)
    }

    colors := map[string]string {
        "AliceBlue":   "#f0f8ff",
        "Coral":       "#ff7F50",
        "DarkGray":    "#a9a9a9",
        "ForestGreen": "#228b22",
    }
    for key, value := range colors {
        fmt.Printf("Key: %s Value: %s\n", key, value)
    }




}
