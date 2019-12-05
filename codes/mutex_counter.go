package main

import (
    "fmt"
    "runtime"
    "sync"
)

var (
    counter int
    wg sync.WaitGroup
    mutex sync.Mutex
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
        mutex.Lock()
        {
            value := counter

            runtime.Gosched()

            value++
            fmt.Printf("%s: %d\n", prefix, value)
            counter = value
        }
        mutex.Unlock()
    }
}
