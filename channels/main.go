package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func calculateSum(ch chan int, n int) {
	defer wg.Done()
	ch <- n * (n + 1) / 2
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide input integer")
		return
	}
	input := os.Args[1]
	n, _ := strconv.Atoi(input)
	ch := make(chan int, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go calculateSum(ch, i)
	}
	wg.Wait()
	close(ch)

	commulativeSum := 0
	for item := range ch {
		commulativeSum += item
	}
	fmt.Println("Supersum of sum of first n numbers", commulativeSum)

}
