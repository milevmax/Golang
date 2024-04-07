package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	//cancel()

	ch := make(chan int)

	go func(ctx context.Context, ch chan int) {
		fmt.Println("start goroutine ...")
		time.Sleep(time.Second * 5)
		fmt.Println("goroutine finishing")
		ch <- 228
	}(ctxTimeout, ch)

	//go func(ch chan int) {
	//	fmt.Println("start goroutine ...")
	//	time.Sleep(time.Second * 2)
	//	fmt.Println("goroutine finishing")
	//	ch <- 228
	//}(ch)

	select {
	case <-ctxTimeout.Done():
		fmt.Println("goroutine finished by timeout")
	case result := <-ch:
		fmt.Println("goroutine sent: ", result)
	}

	fmt.Println("sleep in main...")
	time.Sleep(time.Second * 3)
	fmt.Println("waked up in main!")
}
