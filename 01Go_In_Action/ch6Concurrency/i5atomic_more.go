package main

import (
    "fmt"

    "sync"
    "sync/atomic"
    "time"
)

var (
    // 通知正在执行的gorountine 停止工作的标志
    shutdown int64

    wg sync.WaitGroup
)

func main() {
    wg.Add(2)

    go doWork("A")
    go doWork("B")

    time.Sleep(1 * time.Second)

    //该停止工作
    fmt.Println("Shutown Now!!!!!")
    atomic.StoreInt64(&shutdown, 1)

}

// 检测之前的shutdown标志来决定是否提前终止
func doWork(name string) {
    defer wg.Done()

    for {
        fmt.Printf("Doing %s Work\n", name)
        time.Sleep(250 * time.Millisecond)

        if atomic.LoadInt64(&shutdown) == 1 {
            fmt.Printf("Shutting %s down\n", name)
            break
        }
    }
}