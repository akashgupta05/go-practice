package main

import (
	"fmt"
	"sync"
)

var num = 0
var wg sync.WaitGroup

func sumWorker(mx *sync.Mutex) {
	mx.Lock()
	num++
	mx.Unlock()
	wg.Done()
}

func main() {
	var mx sync.Mutex

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go sumWorker(&mx)
	}

	wg.Wait()
	fmt.Println("Result", num)
}
