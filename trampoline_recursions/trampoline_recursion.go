package main

import (
	"fmt"
	"runtime"
	"time"
)

type trampoline func() (trampoline, int)

func trampolineRecursive(factorial, accumulator int) trampoline {
	if factorial == 0 {
		return func() (trampoline, int) {
			return nil, accumulator
		}
	}

	return func() (trampoline, int) {
		return trampolineRecursive(factorial-1, accumulator*factorial), 0
	}
}

func factorialTrampoline(n int) int {
	accumulator := 1
	fn := trampolineRecursive(n, accumulator)
	for fn != nil {
		fn, accumulator = fn()
	}

	return accumulator
}

func main() {
	var mem runtime.MemStats
	n := 0

	fmt.Println("Please enter the number")
	fmt.Scanf("%d", &n)
	t := time.Now()

	fmt.Println("Factorial result: ", factorialTrampoline(n))

	runtime.ReadMemStats(&mem)
	fmt.Printf("Stacks in use: %v, Time taken: %vms", mem.StackInuse, time.Now().Sub(t).Milliseconds())
}
