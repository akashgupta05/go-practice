package main

import (
	"fmt"
	"time"
)

func func1(ch chan string, str string) {
	ch <- "message from func1 - " + str
}

func func2(ch chan string, str string) {
	ch <- "message from func2 - " + str
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func1(ch1, "something")
	go func2(ch2, "else")

	time.Sleep(2 * time.Second)

	// either output will show
	select {
	case val := <-ch1:
		fmt.Println("Message received :", val)
	case val := <-ch2:
		fmt.Println("Message received :", val)
	default:
		fmt.Println("No Message received")
	}
}
