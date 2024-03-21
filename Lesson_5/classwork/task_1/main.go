package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	wordList := strings.Split(text, " ")
	wordMap := make(map[string]int)
	for _, word := range wordList {
		wordMap[word] += 1
	}
	return wordMap
}

func main() {
	wordsCount := make(map[string]int)
	sampleText := "There are some things that are so unforgivable that they make other things easily forgivable"

	wordsCount = countWords(sampleText)

	fmt.Println(sampleText)
	fmt.Println(wordsCount)

	for k, v := range wordsCount {
		fmt.Println(k, v)
	}
}
