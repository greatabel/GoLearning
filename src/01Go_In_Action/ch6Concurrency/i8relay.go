package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

func main(){
    baton := make(chan int)

    wg.Add(1)

    go Runner(baton)

    baton <- 1

    wg.Wait()

}

func Runner(baton chan int){
    var newRunner int

    // Wait to receive the baton.
    runner := <-baton

    // Start running around the track.
    fmt.Printf("Runner %d Running With Baton\n", runner)

    // New runner to the line.
    if runner != 4 {
        newRunner = runner + 1
        fmt.Printf("Runner %d To The Line\n", newRunner)
        go Runner(baton)
    }

    // Running around the track.
    time.Sleep(100 * time.Millisecond)

    // Is the race over.
    if runner == 4 {
        fmt.Printf("Runner %d Finished, Race Over\n", runner)
        wg.Done()
        return
    }

    // Exchange the baton for the next runner.
    fmt.Printf("Runner %d Exchange With Runner %d\n",
        runner,
        newRunner)

    baton <- newRunner
}
