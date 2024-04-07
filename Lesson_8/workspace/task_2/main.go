package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())

	// Start a goroutine that does some work
	go func(ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine canceled")
				return
			default:
				// Simulate some work
				fmt.Println(i)
				time.Sleep(500 * time.Millisecond) // 0.5 second delay
			}
		}
	}(ctx)

	// Wait a bit and then cancel the context
	time.Sleep(2 * time.Second) // Wait 2 seconds
	cancel()                    // This cancels the context

	// Wait a little for the goroutine to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function done")
}
