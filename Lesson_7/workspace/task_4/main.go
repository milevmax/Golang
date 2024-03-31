package main

import "fmt"

func checkPrime(n int, primeNumbers chan int) bool {
	divider := 2
	for ; divider < n; divider++ {
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

func findAllPrimes(MaxNumber int) []int {
	primeNumbersCh := make(chan int)

	for i := 2; i <= MaxNumber; i++ {
		if checkPrime(i, primeNumbersCh) {

		}
	}
	//primeNumbers := collectPrimeNumbers(MaxNumber, primeNumbersCh)

	close(primeNumbersCh)
	return []int{}
}

func main() {
	p := findAllPrimes(1000)
	fmt.Println(p)
}

// чи можливо створити канал з типо слайс і взагалі які типи в каналі можна використовувати?
// як обмежити кількість активних каналів?
