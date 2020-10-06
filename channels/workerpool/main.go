package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// worker that makes cubes
func cubeWorker(tasks, results chan int, i int) {
	for num := range tasks {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Sending result by worker :", i)
		results <- num * num * num
	}

	wg.Done()
}

func main() {
	fmt.Println("Start")
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// 2 workers goroutines
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go cubeWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2
	}

	fmt.Println("Assigned 5 tasks")

	// closing tasks
	close(tasks)

	// wait until worker finshes the job
	wg.Wait()

	//closing results
	close(results)

	for result := range results {
		fmt.Println("Result :", result)
	}

	fmt.Println("Finish")
}
