package main

import (
	"fmt"
	"math/rand/v2"
)

func RandomGenerator(countNumbers int, randomChan chan int) {
	var rn int
	for i := 0; i < countNumbers; i++ {
		rn = rand.IntN(1000)
		randomChan <- rn
	}
	close(randomChan)
}

func Average(randomChan chan int, meanChan chan float64) {
	var sum, count int
	for v := range randomChan {
		sum += v
		count += 1
	}
	meanChan <- float64(sum) / float64(count)
	close(meanChan)
}

func PrintAverage(meanChan chan float64, done chan struct{}) {
	avgVal := <-meanChan
	fmt.Printf("mean value: %f", avgVal)
	done <- struct{}{}
}

func main() {

	countNumbers := 10
	randomChan := make(chan int)
	meanChan := make(chan float64)
	done := make(chan struct{})

	go RandomGenerator(countNumbers, randomChan)
	go Average(randomChan, meanChan)
	go PrintAverage(meanChan, done)
	<-done
}
