package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	Goroutines = 4
	taskLoad   = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(Goroutines)
	for i := 1; i <= Goroutines; i++ {
		go worker(tasks, i)
	}

	for i := 1; i <= taskLoad; i++ {
		tasks <- fmt.Sprintf("Task: %d", i)
	}

	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d Started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d Finished %s\n", worker, task)
	}
}
