package main

import (
	"fmt"
	"strings"
)

var (
	sourceText_1 string = "Stars are massive celestial bodies made of hydrogen and helium. They emit light and heat from nuclear reactions at their cores. The closest star to Earth is the Sun, which is the center of our solar system. Stars vary in size, color, and brightness, from small red dwarfs to massive blue supergiants. Over time, stars go through life cycles, being born from nebulae and eventually dying, sometimes as supernovae. The number of stars in the universe is vast, with billions in our Milky Way galaxy alone. Observing stars has helped astronomers understand the cosmos and our place within it."
	sourceText_2 string = "Our solar system is home to eight planets, including Earth, orbiting around a single star, the Sun. The Milky Way, a vast galaxy containing our solar system, spans across billions of stars, each with their potential systems. Stars, including our Sun, primarily consist of hydrogen and helium, the two simplest and most abundant elements in the universe. The formation of stars begins in nebulae, where clouds of hydrogen and helium collapse under gravity to ignite nuclear fusion."
)

func processText(sentences []string) []string {
	for i, v := range sentences {
		senIter := strings.ToLower(v)
		senIter = strings.TrimSpace(senIter)
		sentences[i] = senIter
	}
	return sentences
}

func findSentence(targetSubStr string, sentences []string) (sentencesSlice []string) {

	sentencesSlice = make([]string, 0, len(sentences))
	var lenSubStr int = len(targetSubStr)

	for _, sen := range sentences {
		for i := 0; i < len(sen)-lenSubStr; i++ {
			IterSubStr := sen[i : i+lenSubStr]

			if IterSubStr == targetSubStr {
				sentencesSlice = append(sentencesSlice, sen)
				break
			}
		}
	}
	return sentencesSlice
}

func main() {
	sourceText := sourceText_1 + sourceText_2
	textSentences := strings.Split(sourceText, ".")
	processesText := processText(textSentences)
	fmt.Print(processesText)

	fmt.Println()
	fmt.Println()

	testStr := "12345abcde"
	testStrCut := testStr[2:6]
	fmt.Print(testStrCut)

	fmt.Println()
	fmt.Println()

	targetStr := "solar system"
	resultSentences := findSentence(targetStr, processesText)
	for _, s := range resultSentences {
		fmt.Println(s)
	}
}
