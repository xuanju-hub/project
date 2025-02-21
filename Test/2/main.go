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
	// 创建一个用于接收工作结果的通道
	workDone := make(chan int)
	// 启动一个goroutine来执行耗时操作，并将结果发送到workDone通道
	go func() {
		// 调用LongWork函数执行耗时操作，并将结果发送到workDone通道
		workDone <- LongWork()
	}()

	// 使用select语句等待工作结果或超时
	select {
	case n := <-workDone:
		// 如果工作完成，从workDone通道接收结果并打印
		fmt.Println(n)
		// 返回工作结果
		return n
	case <-time.After(timeout):
		// 如果超时，打印"timeout"
		fmt.Println("timeout")
		// 返回0表示超时
		return 0
	}
}

func main() {
	Handle()
}
