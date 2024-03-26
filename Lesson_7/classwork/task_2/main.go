package main

import "fmt"

func sendNumber(ch chan int) {
	for i := 0; i < 1000000; i += 1 + i*i {
		ch <- i
	}
	close(ch)
}

func printNumbers(ch chan int, d chan struct{}) {
	for i := range ch {
		fmt.Println("get value: ", i)
	}
	fmt.Println("print done!")
	d <- struct{}{}
}

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	go sendNumber(ch)
	go printNumbers(ch, done)
	<-done
}
