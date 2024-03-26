package main

import (
	"fmt"
	"math/rand/v2"
)

type Stats struct {
	minVal int
	maxVal int
}

func (s Stats) String() string {
	return fmt.Sprintf("min value: %d\nmax value: %d", s.minVal, s.maxVal)
}

func RandomGenerator(countNumbers, maxPossibleValue int, randomChan chan int, statsChan chan Stats, done chan struct{}) {
	var rn int
	for i := 0; i < countNumbers; i++ {
		rn = rand.IntN(maxPossibleValue)
		randomChan <- rn
	}
	close(randomChan)

	resultStats := <-statsChan
	fmt.Println(resultStats)
	done <- struct{}{}
}

func MinMax(randomChan chan int, statsChan chan Stats) {
	randVal := <-randomChan
	var minValue, maxValue int = randVal, randVal
	for v := range randomChan {
		if v > maxValue {
			maxValue = v
		}

		if v < minValue {
			minValue = v
		}
	}
	resultStats := Stats{minValue, maxValue}
	statsChan <- resultStats
	close(statsChan)
}

func main() {
	var (
		countNumbers     int = 5
		maxPossibleValue int = 1000
	)
	randomChan := make(chan int)
	statsChan := make(chan Stats)
	done := make(chan struct{})

	go RandomGenerator(countNumbers, maxPossibleValue, randomChan, statsChan, done)
	go MinMax(randomChan, statsChan)
	<-done
}
