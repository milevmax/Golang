package main

import "fmt"

func countTwoCharacters(text string, char1, char2 rune) (count1, count2 int) {

	for _, char := range text {
		switch char {
		case char1:
			count1 += 1
		case char2:
			count2 += 1
		}
	}
	return
}

func main() {
	char1, char2 := 'a', 'c'
	text := "abcd abc cd"
	count1, count2 := countTwoCharacters(text, char1, char2)
	fmt.Printf("character `%c` occurred %d times, character `%c` occurred %d times\n", char1, count1, char2, count2)
}
