package main

import "fmt"

func checkPrime(n int, primeNumbers chan int) {
	divider := 2
	for ; divider < n; divider++ {
		if n%divider == 0 {
			break
		}
	}
	if divider == n {
		primeNumbers <- n
	}
}

func collectPrimeNumbers(outputLen int, primeNumbersCh chan int) (result []int) {
	result = make([]int, 0, outputLen)
	for pn := range primeNumbersCh {
		result = append(result, pn)
	}
	return result
}

func findAllPrimes(MaxNumber int) []int {
	primeNumbersCh := make(chan int)

	for i := 2; i <= MaxNumber; i++ {
		go checkPrime(i, primeNumbersCh)
	}
	primeNumbers := collectPrimeNumbers(MaxNumber, primeNumbersCh)

	close(primeNumbersCh)
	return primeNumbers
}

func main() {
	p := findAllPrimes(1000)
	fmt.Println(p)
}

// чи можливо створити канал з типо слайс і взагалі які типи в каналі можна використовувати?
// як обмежити кількість активних каналів?
