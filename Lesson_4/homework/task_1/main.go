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
	sentencesProcessed := make([]string, len(sentences))
	copy(sentencesProcessed, sentences)
	for i, v := range sentencesProcessed {
		senIter := strings.ToLower(v)
		senIter = strings.TrimSpace(senIter)
		sentencesProcessed[i] = senIter
	}
	return sentencesProcessed
}

func findSentence(targetSubStr string, sentences []string) (sentencesIndexes []int) {

	sentences = processText(sentences)
	targetSubStr = strings.ToLower(targetSubStr)

	sentencesIndexes = make([]int, 0, len(sentences))
	var lenSubStr int = len(targetSubStr)

	for ind, sen := range sentences {
		for i := 0; i < len(sen)-lenSubStr; i++ {
			IterSubStr := sen[i : i+lenSubStr]

			if IterSubStr == targetSubStr {
				sentencesIndexes = append(sentencesIndexes, ind)
				break
			}
		}
	}
	return sentencesIndexes
}

func getSenByIdexes(sentences []string, sentencesIndexes []int) (sliceSentences []string) {

	sliceSentences = make([]string, 0, len(sentencesIndexes))

	for _, i := range sentencesIndexes {
		sliceSentences = append(sliceSentences, sentences[i])
	}
	return sliceSentences
}

func findMatches(text string, target string) (Sentences []string) {

	textSentences := strings.Split(text, ".")
	resultIndexes := findSentence(target, textSentences)
	resultSentences := getSenByIdexes(textSentences, resultIndexes)
	return resultSentences
}

func main() {
	sourceText := sourceText_1 + sourceText_2
	targetstr := "milky way"

	resultSentences := findMatches(sourceText, targetstr)
	fmt.Printf("\nTarget string: %s\n", targetstr)
	fmt.Println("Matched sentences:")
	for _, s := range resultSentences {
		fmt.Println(s)
	}
}
