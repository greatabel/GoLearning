package pool

import (
    "errors"
    "io"
    "log"
    "sync"
)

type Pool struct {
    m         sync.Mutex
    resources chan io.Closer
    factory   func() (io.Closer, error)
    closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
    if size <= 0 {
        return nil, errors.New("Size value too small.")
    }

    return &Pool{
        factory:   fn,
        resources: make(chan io.Closer, size),
    }, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
    select {
    // Check for a free resource.
    case r, ok := <-p.resources:
        log.Println("Acquire:", "Shared Resource")
        if !ok {
            return nil, ErrPoolClosed
        }
        return r, nil

    // Provide a new resource since there are none available.
    default:
        log.Println("Acquire:", "New Resource")
        return p.factory()
    }
}

func (p *Pool) Release(r io.Closer) {
    // Secure this operation with the Close operation.
    p.m.Lock()
    defer p.m.Unlock()

    // If the pool is closed, discard the resource.
    if p.closed {
        r.Close()
        return
    }

    select {
    // Attempt to place the new resource on the queue.
    case p.resources <- r:
        log.Println("Release:", "In Queue")

    // If the queue is already at cap we close the resource.
    default:
        log.Println("Release:", "Closing")
        r.Close()
    }
}

func (p *Pool) Close() {
    // Secure this operation with the Release operation.
    p.m.Lock()
    defer p.m.Unlock()

    // If the pool is already close, don't do anything.
    if p.closed {
        return
    }

    // Set the pool as closed.
    p.closed = true

    // Close the channel before we drain the channel of its
    // resources. If we don't do this, we will have a deadlock.
    close(p.resources)

    // Close the resources
    for r := range p.resources {
        r.Close()
    }
}