package main

import "fmt"

func sendNumber(ch chan int) {
	for i := 0; i < 1000000; i += 1 + i*i {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go sendNumber(ch)

	//for val := range ch {
	//	fmt.Println(val)
	//}

	for {
		val, open := <-ch
		if open {
			fmt.Println(val)
		} else {
			break
		}
	}

	fmt.Println("Done!")
}
