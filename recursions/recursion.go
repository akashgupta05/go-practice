package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	var mem runtime.MemStats
	n := 0

	fmt.Println("Please enter the number")
	fmt.Scanf("%d", &n)
	t := time.Now()

	fmt.Println("Factorial result: ", factorial(n))

	runtime.ReadMemStats(&mem)
	fmt.Printf("Stacks in use: %v, Time taken: %vms", mem.StackInuse, time.Now().Sub(t).Milliseconds())
}
