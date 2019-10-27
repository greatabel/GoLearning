package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

var (
    counter int64

    wg sync.WaitGroup
)

func main() {
    wg.Add(2)

    go incCounter(1)
    go incCounter(2)

    wg.Wait()

    fmt.Println("Counter:", counter)

}

func incCounter(id int) {
    defer wg.Done()

    for count := 0; count < 2; count++ {
        // 安全地对counter 加1
        atomic.AddInt64(&counter, 1)

        runtime.Gosched()
    }
}