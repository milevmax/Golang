package main

import "fmt"

func sendNumber(ch chan int) {
	//for i := 0; i < 1000000; i += 1 + i*i {
	//	ch <- i
	//}
	//close(ch)

	ch <- 10
	ch <- 20
	ch <- 30
}

func printNumber(ch chan int, wait chan bool) {
	for val := range ch {
		fmt.Println(val)
	}
	wait <- true
	close(wait)
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go sendNumber(ch)
	go printNumber(ch, done)
	<-done
	fmt.Println("Done!")
}
