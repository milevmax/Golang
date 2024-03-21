package main

import (
	"fmt"
	"homework5/Util"
	"strings"
)

type TextRedactor struct {
	text      string
	sentences []string
	indexMap  map[string][]int
}

func (t *TextRedactor) Init(text string) {
	t.text = text
	t.sentences = Util.TextToSlice(text)
	t.indexMap = make(map[string][]int)
}

func (t *TextRedactor) IndexText() {
	sentences := Util.ProcessSentences(t.sentences)
	splitSentences := Util.SplitSentences(sentences)

	lastSeenIn := make(map[string]int)
	indexMap := make(map[string][]int)

	for senIndex, sen := range splitSentences {
		for _, word := range sen {

			senIndxLast, ok := lastSeenIn[word]

			if senIndex > senIndxLast {
				indexMap[word] = append(indexMap[word], senIndex)
				lastSeenIn[word] = senIndex
			}

			if !ok && senIndex == 0 {
				indexMap[word] = append(indexMap[word], senIndex)
				lastSeenIn[word] = senIndex
			}
		}
	}
	t.indexMap = indexMap
}

func (t *TextRedactor) Search(target string) (sliceSentences []string) {
	target = strings.ToLower(target)
	targetIndexes := t.indexMap[target]
	sliceSentences = Util.GetSentencesByIndexes(t.sentences, targetIndexes)
	return sliceSentences
}

var (
	sourceText_1 string = "Stars are massive celestial bodies made of hydrogen and helium. They emit light and heat from nuclear reactions at their cores. The closest star to Earth is the Sun, which is the center of our solar system. Stars vary in size, color, and brightness, from small red dwarfs to massive blue supergiants. Over time, stars go through life cycles, being born from nebulae and eventually dying, sometimes as supernovae. The number of stars in the universe is vast, with billions in our Milky Way galaxy alone. Observing stars has helped astronomers understand the cosmos and our place within it."
	sourceText_2 string = "Our solar system is home to eight planets, including Earth, orbiting around a single star, the Sun. The Milky Way, a vast galaxy containing our solar system, spans across billions of stars, each with their potential systems. Stars, including our Sun, primarily consist of hydrogen and helium, the two simplest and most abundant elements in the universe. The formation of stars begins in nebulae, where clouds of hydrogen and helium collapse under gravity to ignite nuclear fusion."
)

func main() {
	sourceText := sourceText_1 + sourceText_2

	MyRedactor := TextRedactor{}
	MyRedactor.Init(sourceText)
	MyRedactor.IndexText()

	targetStr := "stars"
	resultSentences := MyRedactor.Search(targetStr)

	fmt.Printf("\nTarget string: %s\n", targetStr)
	fmt.Println("Matched sentences:")
	for _, s := range resultSentences {
		fmt.Println(s)
	}
}
