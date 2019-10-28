package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    numberGoroutines = 4  // 要使用的goroutine的数量
    taskLoad         = 10 // 要处理的工作的数量
)

var wg sync.WaitGroup

//init初始化包，Go语言运行时会在其他代码执行之前
func init() {
    // Seed the random number generator.
    rand.Seed(time.Now().Unix())
}

func main(){
    // 创建一个有缓冲的通道来管理工作
    tasks := make(chan string, taskLoad)

    wg.Add(numberGoroutines)

    for gr := 1; gr <= numberGoroutines; gr++ {
        go worker(tasks, gr)
    }

    for post := 1; post <= taskLoad; post++ {
        tasks <- fmt.Sprintf("任务: %d", post)
    }
    close(tasks)
    wg.Wait()

}

// 同有缓冲的通道传入工作
func worker(tasks chan string ,worker int) {
    defer wg.Done()

    for {
        // 等待分配工作
        task, ok := <- tasks
        if !ok {
            // 通道已经空了，并且已被关闭
            fmt.Printf("Worker: %d : 关闭ing\n", worker)
            return
        }

        // 开始工作
        fmt.Printf("Worker: %d : 开始干 %s\n", worker, task)

        // 随机一段时间模拟工作
        sleep := rand.Int63n(100)
        time.Sleep(time.Duration(sleep) * time.Millisecond)

        // 工作完成
        fmt.Printf("Worker: %d : Completed %s\n", worker, task)
    }
}