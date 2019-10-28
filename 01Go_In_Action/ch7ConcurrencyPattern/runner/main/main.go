package main

import (
    "log"
    "os"
    "time"

    "01Go_In_Action/ch7ConcurrencyPattern/runner"
)

// timeout is the number of second the program has to finish.
const timeout = 3 * time.Second

func main() {
    log.Println("Starting work.")

    // Create a new timer value for this run.
    r := runner.New(timeout)

    // Add the tasks to be run.
    r.Add(createTask(), createTask(), createTask())

    // Run the tasks and handle the result.
    if err := r.Start(); err != nil {
        switch err {
        case runner.ErrTimeout:
            log.Println("Terminating due to timeout.")
            os.Exit(1)
        case runner.ErrInterrupt:
            log.Println("Terminating due to interrupt.")
            os.Exit(2)
        }
    }

    log.Println("Process ended.")
}

func createTask() func(int) {
    return func(id int) {
        log.Printf("处理器 -任务 #%d.", id)
        time.Sleep(time.Duration(id) * time.Second)
    }
}