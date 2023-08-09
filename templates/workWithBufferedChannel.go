package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg2 sync.WaitGroup

func init() {
	rand.NewSource(time.Now().Unix())
}

func worker(tasks chan string, worker int) {
	defer wg2.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

func main() {
	tasks := make(chan string, taskLoad)

	wg2.Add(numberGoroutines)

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	close(tasks)

	wg2.Wait()
}
