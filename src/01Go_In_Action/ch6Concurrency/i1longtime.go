package main

import (
    "fmt"
    "runtime"
    "sync"
)

var wg sync.WaitGroup

func main() {
    runtime.GOMAXPROCS(1)

    wg.Add(2)

    fmt.Println("构建Goroutines")
    go printPrime("A")
    go printPrime("B")

    fmt.Println("等待完成")
    wg.Wait()
    fmt.Println("终止程序")  
}

func printPrime(prefix string) {
    defer wg.Done()

next:
    for outer := 2; outer < 5000; outer++ {
        for inner := 2; inner < outer; inner++ {
            if outer % inner == 0 {
                continue next
            }
        }
        fmt.Printf("%s:%d\n", prefix, outer)
    }
    fmt.Println("完成", prefix)
}