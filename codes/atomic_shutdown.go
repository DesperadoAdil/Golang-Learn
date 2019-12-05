package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

var (
    shutdown int64
    wg sync.WaitGroup
)

func main()  {
    wg.Add(2)

    fmt.Println("Start Goroutines!")
    go doWork("A")
    go doWork("B")

    time.Sleep(1 * time.Second)

    fmt.Println("Shutdown Now")
    atomic.StoreInt64(&shutdown, 1)
    wg.Wait()
    fmt.Println("Terminating Program")
}

func doWork(prefix string)  {
    defer wg.Done()

    for {
        fmt.Printf("Doing %s Work\n", prefix)
        time.Sleep(250 * time.Millisecond)

        if atomic.LoadInt64(&shutdown) == 1 {
            fmt.Printf("Shutting %s Work\n", prefix)
            break
        }
    }
}
