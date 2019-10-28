package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "01Go_In_Action/ch3Package_toolchain/words"
)

func main() {
    filename := os.Args[1]

    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        return
    }

    text := string(contents)

    count := words.CountWords(text)
    fmt.Printf("There are %d words in your text. \n", count)
}