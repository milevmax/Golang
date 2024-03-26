package main

import (
	"fmt"
	"time"
)

func checkPrime(n int) bool {
	divider := 2
	for ; divider < n; divider++ {
		time.Sleep(20 * time.Millisecond)
		if n%divider == 0 {
			break
		}
	}
	if divider == n {
		return true
	} else {
		return false
	}
}

func findAllPrimes(N int, primeNumbers chan int) {
	for i := 2; i <= N; i++ {
		if checkPrime(i) {
			primeNumbers <- i
		}
	}
	close(primeNumbers)
}

func abobaMessage(ch chan string) {
	ch <- "aboba"
	//close(ch)
}

func abobaReciever(ch chan int, done chan struct{}) {
	for v := range ch {
		time.Sleep(20 * time.Millisecond)
		fmt.Println("aboba recieved: ", v)
		//if v == 222 {
		//	break
		//}
	}
	fmt.Println("loop done")
	done <- struct{}{}
	//ch <- 999
}

func main() {
	//c := make(chan int)
	//maxPrime := 23
	//go findAllPrimes(maxPrime, c)
	//
	//for pn := range c {
	//	fmt.Println(pn)
	//}
	//fmt.Println("Done!")

	//------------------------------------------
	//ch := make(chan string)
	//
	//go func() {
	//	ch <- "aboba"
	//}()
	//
	//a := <-ch
	//fmt.Printf("got: %q\n", a)
	//------------------------------------------
	//ch := make(chan string)
	//go abobaMessage(ch)
	//aboba := <-ch
	//fmt.Printf("got: %q\n", aboba)
	//close(ch)
	//anotherAboba := <-ch
	//fmt.Printf("got: %q\n", anotherAboba)
	//anotherAboba, open := <-ch
	//fmt.Println(anotherAboba, open)
	//------------------------------------------
	ch2 := make(chan int)
	done := make(chan struct{})
	go abobaReciever(ch2, done)

	ch2 <- 12
	ch2 <- 10
	ch2 <- 222
	//<-ch2

	close(ch2)
	<-done
}
