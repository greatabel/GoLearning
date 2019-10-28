package main

import "fmt"

// duration is a type with a base type of int.
type duration int

// format pretty-prints the duration value.
func (d *duration) pretty() string {
    return fmt.Sprintf("Duration: %d", *d)
}

// main is the entry point for the application.
func main() {
    // duration(42).pretty()
    var dur duration
    dur = 100
    
    fmt.Println(dur.pretty())


    // ./listing46.go:17: cannot call pointer method on duration(42)
    // ./listing46.go:17: cannot take the address of duration(42)
}