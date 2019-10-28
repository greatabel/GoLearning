package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    fmt.Println(runtime.NumCPU())
    // 这个函数允许程序 更改调度器可以使用的逻辑处理器的数量
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("开始 Goroutines")

    go func() {
        defer wg.Done()

        for count :=0; count < 3; count++ {
            for char := 'a'; char < 'a'+26; char++ {
                fmt.Printf("%c ", char)
            }
            fmt.Println()
        }
        
    }()

    go func() {
        defer wg.Done()
        for count :=0; count < 3; count++ {
            for char := 'A'; char < 'A'+26; char++ {
                fmt.Printf("%c ", char)
            }
            fmt.Println()
        }
        
    }()

    fmt.Println("结束等待")
    wg.Wait()

    fmt.Println("\n终止程序ing")
}