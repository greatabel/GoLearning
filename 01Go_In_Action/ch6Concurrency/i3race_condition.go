package main

import (
    "fmt"
    "runtime"
    "sync"
)

var (
    // counter是所有goroutine都要增加其值的变量
    counter int

    wg sync.WaitGroup
)

func main() {
    wg.Add(2)
    go incCounter(1)
    go incCounter(2)

    wg.Wait()
    fmt.Println("最终 Counter:", counter)


}

func incCounter(id int){
    defer wg.Done()

    for count := 0; count < 2; count++ {
        // 捕获counter 值
        value := counter
        // 当前goroutine从线程退出，并放回队列
        runtime.Gosched()

        value++
        counter = value
    }
}