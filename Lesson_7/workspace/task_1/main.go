package main

import (
	"fmt"
	"time"
)

func findPrime(N int, primeNumbers chan int) {
	for i := 2; i <= N; i++ {
		divider := 2
		for ; divider <= i; divider++ {
			if i%divider == 0 {
				break
			}
		}
		if divider == i {
			primeNumbers <- i
		}
	}
	close(primeNumbers)
}

func main() {
	c := make(chan int)
	maxPrime := 23

	go findPrime(maxPrime, c)

	for pn := range c {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(pn)
	}

	fmt.Println("Done!")
}
