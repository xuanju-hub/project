package main

import (
	"fmt"
	"time"
)

const timeout = 300 * time.Millisecond

func LongWork() int {
	time.Sleep(200 * time.Millisecond)
	return 1
}

func Handle() int {
	workDone := make(chan int)
	go func() {
		workDone <- LongWork()
	}()

	select {
	case n := <-workDone:
		fmt.Println(n)
		return n
	case <-time.After(timeout):
		fmt.Println("timeout")
		return 0
	}
}

func main() {

}
