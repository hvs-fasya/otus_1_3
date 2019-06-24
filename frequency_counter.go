package otus_1_3

import (
	"strings"
	"unicode"
)

var punctuation = []string{`.`, `,`, `:`, `;`, `'`, `"`, `!`, `?`, `(`, `)`}

//MostFrequent return 'max' most frequent words in 'text'
func MostFrequent(text string, max int) []string {
	var res = make([]string, 0)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	var wordFrequencyMap = make(map[string]int, 0) // map["word1": 2, "word2": 1]
	var maxCount int
	for _, w := range words {
		var lowered = strings.ToLower(w)
		if _, ok := wordFrequencyMap[lowered]; !ok {
			wordFrequencyMap[lowered] = 0
		}
		wordFrequencyMap[lowered]++
		if wordFrequencyMap[lowered] > maxCount {
			maxCount = wordFrequencyMap[lowered]
		}
	}
	var frequenciesArray = make([][]string, maxCount+1, maxCount+1) // [[] [word1 word2] [word3 word4]]
	for word, count := range wordFrequencyMap {
		frequenciesArray[count] = append(frequenciesArray[count], word)
	}
	for i := len(frequenciesArray) - 1; i > 0; i-- {
		for _, word := range frequenciesArray[i] {
			if len(res) >= max {
				return res
			}
			res = append(res, word)
		}
	}
	return res
}
