package Util

import (
	"regexp"
	"strings"
)

func TextToSlice(text string) []string {
	pattern := "[.?!]"
	result := regexp.MustCompile(pattern).Split(text, -1)
	return result
}

func ProcessSentences(sentences []string) []string {
	sentencesProcessed := make([]string, len(sentences))
	for i, v := range sentences {
		senIter := strings.ToLower(v)
		senIter = strings.TrimSpace(senIter)
		sentencesProcessed[i] = senIter
	}
	return sentencesProcessed
}

func GetSentencesByIndexes(sentences []string, sentencesIndexes []int) (sliceSentences []string) {

	sliceSentences = make([]string, len(sentencesIndexes))

	for si, i := range sentencesIndexes {
		sliceSentences[si] = sentences[i]
	}
	return sliceSentences
}

func SentenceToSlice(sentence string) []string {
	words := strings.Split(sentence, " ")
	re, _ := regexp.Compile(`[^\w]`)
	for i, word := range words {
		word = strings.ToLower(word)
		word = re.ReplaceAllString(word, "")
		words[i] = word
	}
	return words
}

func SplitSentences(sentences []string) (output [][]string) {

	output = make([][]string, len(sentences))

	for i := range output {
		output[i] = make([]string, 20)
	}

	for i, sentence := range sentences {
		words := SentenceToSlice(sentence)
		output[i] = words
	}
	return output
}
