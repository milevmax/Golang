package main

import (
	"context"
	"fmt"
	"time"
)

func runStoppintWithContext(ctx context.Context, tasks chan int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("got stop signal")
			close(tasks)
			return
		default:
			tasks <- i
		}
	}
}

func runContextCancelFunc(contextCancelFunc context.CancelFunc) {
	time.Sleep(5 * time.Second)
	fmt.Println("Stop printing signal")
	contextCancelFunc()
}
func main() {

	ctx, contextCancelFunc := context.WithCancel(context.Background())
	go runContextCancelFunc(contextCancelFunc)

	tasks := make(chan int)
	go runStoppintWithContext(ctx, tasks)

	for task := range tasks {
		fmt.Printf("task %d started\n", task)
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("task %d Done!\n", task)
	}
}
