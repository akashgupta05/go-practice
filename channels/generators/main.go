package main

import (
	"fmt"
)

// generates n fibonacci numbers
func generateFib(n int) <-chan int {
	c := make(chan int, n)

	// run fibonacci generator concurrently
	go func() {
		for i, j, k := 0, 1, 0; k < n; i, j, k = i+j, i, k+1 {
			c <- i
		}
		close(c)
	}()
	return c
}

func main() {
	// 10 fibonacci numbers from channel
	for num := range generateFib(10) {
		fmt.Println("Current Fibonacci number:", num)
	}
}
