package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    //创建无缓冲通道
    court := make(chan int)

    wg.Add(2)

    go player("abel", court)
    go player("chuchu", court)

    // 发球
    court <- 1

    wg.Wait()
}

// 模拟一个球手
func player(name string, court chan int) {
    defer wg.Done()
    for {
        // 等待球被击打过来
        ball, ok := <-court
        if !ok {
            // If the channel was closed we won.
            // 通道被关闭，我们就赢了
            fmt.Printf("Player %s 赢了\n", name)
            return
        }

        // Pick a random number and see if we miss the ball.
        n := rand.Intn(100)
        if n%13 == 0 {
            fmt.Printf("Player %s Missed 失手了\n", name)

            // 关闭通道，表示我们输掉了
            close(court)
            return
        }

        // 显示击球数目，并将击球数加1
        fmt.Printf("Player %s 击球 %d 次\n", name, ball)
        ball++

        // 把球打给对手
        court <- ball
    }
}