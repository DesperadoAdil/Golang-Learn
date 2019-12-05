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

func main()  {
    wg.Add(2)

    fmt.Println("Start Goroutines!")
    go increaseCounter("A")
    go increaseCounter("B")

    fmt.Println("Waiting To Finish")
    wg.Wait()
    fmt.Println("Terminating Program")
    fmt.Printf("Final Counter: %d\n", counter)
}

func increaseCounter(prefix string)  {
    defer wg.Done()

    for count := 0; count < 2; count++ {
        atomic.AddInt64(&counter, 1)
        runtime.Gosched()
    }
}
