package main

import (
	"fmt"
	"regexp"
	"strings"
)

var myArr [10]int8

var mySlice []int16

var myMap map[int]string

func main() {

	arr := []string{"Hash map data structures use a hash function, which turns a key into an index within an underlying array.",
		"The hash function can be used to access an index when inserting a value or retrieving a value from a hash map.",
		"Hash map underlying data structure", "Hash maps are built on top of an underlying array data structure using an indexing system.",
		"Each index in the array can store one key-value pair.",
		"If the hash map is implemented using chaining for collision resolution, each index can store another data structure such as a linked list, which stores all values for multiple keys that hash to the same index.",
		"Each Hash Map key can be paired with only one value. However, different keys can be paired with the same value."}

	//fmt.Println(arr)
	//
	//fmt.Print(myMap)
	//myMap = make(map[int]string)
	//myMap[0] = "123"
	//fmt.Print(myMap)
	//
	//r := SentenceToSlice(arr[0])
	//for _, v := range r {
	//	fmt.Println(v)
	//}
	//
	//
	r := SplitSentences(arr)
	fmt.Println(r)

	res := updateWordsMap(arr)
	fmt.Println(res)

	//m := make(map[string]int)
	//m["k1"] = 7
	//m["k2"] = 13
	//fmt.Println("map:", m)
	//
	//v1, ok1 := m["k1"]
	//fmt.Println(v1, ok1)
	//v3, ok3 := m["k3"]
	//fmt.Println(v3, ok3)

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
	//output := make([7][]string, 7)

	output = make([][]string, 7)

	for i := range output {
		output[i] = make([]string, 10) // Each inner slice has a length of 3 elements
	}

	for i, sentence := range sentences {
		words := SentenceToSlice(sentence)
		output[i] = words
	}
	return output
}

func updateWordsMap(sentences []string) map[string][]int {

	lastSeenIn := make(map[string]int)
	output := make(map[string][]int)

	//var wordString string

	splitedSentences := SplitSentences(sentences)

	for senIndx, sen := range splitedSentences {
		for _, word := range sen {

			senIndxLast, ok := lastSeenIn[word]

			if senIndx > senIndxLast {
				output[word] = append(output[word], senIndx)
				lastSeenIn[word] = senIndx
			}

			if !ok && senIndx == 0 {
				output[word] = append(output[word], senIndx)
				lastSeenIn[word] = senIndx
			}
		}
	}
	return output
}
