package main

import (
    "fmt"
    "runtime"
    "sync"
)

var (
    // counter is a variable incremented by all goroutines.
    counter int

    // wg is used to wait for the program to finish.
    wg sync.WaitGroup

    // // mutex 用来定义一段代码临界区
    mutex sync.Mutex
)

func main() {
    // Add a count of two, one for each goroutine.
    wg.Add(2)

    // Create two goroutines.
    go incCounter(1)
    go incCounter(2)

    // Wait for the goroutines to finish.
    wg.Wait()
    fmt.Printf("Final Counter: %d\n", counter)
}

func incCounter(id int){
    defer wg.Done()

    for count := 0; count < 2; count++ {
        // Only allow one goroutine through this
        // critical section at a time.
        mutex.Lock()
        {
            // Capture the value of counter.
            value := counter

            // Yield the thread and be placed back in queue.
            runtime.Gosched()

            // Increment our local value of counter.
            value++

            // Store the value back into counter.
            counter = value
        }
        mutex.Unlock()
        // Release the lock and allow any
        // waiting goroutine through.
    }
}