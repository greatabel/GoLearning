// Example is provided with help by Gabriel Aszalos.
// runner 包管理器：处理任务的运行和生命周期
package runner

import (
    "errors"
    "os"
    "os/signal"
    "time"
)

// Runner在给定的时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
    // interrupt通道 报告从操作系统发送的信号
    interrupt chan os.Signal

    // complete通道 报告处理任务已经完成
    complete chan error

    // timeout 报告处理任务已经超时
    timeout <-chan time.Time

    // tasks持有一组以索引顺序依次执行的函数
    tasks []func(int)
}

// ErrTimeout is returned when a value is received on the timeout channel.
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt is returned when an event from the OS is received.
var ErrInterrupt = errors.New("received interrupt")

// New returns a new ready-to-use Runner.
func New(d time.Duration) *Runner {
    return &Runner{
        interrupt: make(chan os.Signal, 1),
        complete:  make(chan error),
        timeout:   time.After(d),
    }
}

// 将一个任务附加到Runner上。任务是一个接受一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
    r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务，并监视通道事件
func (r *Runner) Start() error {
    // We want to receive all interrupt based signals.
    signal.Notify(r.interrupt, os.Interrupt)

    // Run the different tasks on a different goroutine.
    go func() {
        r.complete <- r.run()
    }()

    select {
    // Signaled when processing is done.
    case err := <-r.complete:
        return err

    // Signaled when we run out of time.
    case <-r.timeout:
        return ErrTimeout
    }
}

// run executes each registered task.
func (r *Runner) run() error {
    for id, task := range r.tasks {
        // Check for an interrupt signal from the OS.
        if r.gotInterrupt() {
            return ErrInterrupt
        }

        // Execute the registered task.
        task(id)
    }

    return nil
}

// gotInterrupt verifies if the interrupt signal has been issued.
func (r *Runner) gotInterrupt() bool {
    select {
    // Signaled when an interrupt event is sent.
    case <-r.interrupt:
        // Stop receiving any further signals.
        signal.Stop(r.interrupt)
        return true

    // Continue running as normal.
    default:
        return false
    }
}
