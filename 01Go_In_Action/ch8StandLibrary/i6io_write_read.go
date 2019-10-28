package main

import (
    "bytes"
    "fmt"
    "os"
)

func main() {
    var b bytes.Buffer

    b.Write([]byte("Hello "))

    fmt.Fprintf(&b, "World from 阿宝!\n")

    b.WriteTo(os.Stdout)
}