package otus_1_3

import (
	"math/rand"
	"sort"
	"strings"
	"unicode"
)

type wordCount struct {
	word  string
	count int
}

//MostFrequent return 'max' most frequent words in 'text' - use sort by counting
func MostFrequent(text string, max int) []string {
	var res = make([]string, 0)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	var wordFrequencyMap = make(map[string]int, 0) // map["word1": 2, "word2": 1]
	var maxCount int
	for _, w := range words {
		var lowered = strings.ToLower(w)
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

//MostFrequentWithSliceSort return 'max' most frequent words in 'text' - use slice sort for most frequent search
func MostFrequentWithSliceSort(text string, max int) []string {
	var res = make([]string, 0)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	var wordFrequencyMap = make(map[string]int, 0) // map["word1": 2, "word2": 1]
	var maxCount int
	for _, w := range words {
		var lowered = strings.ToLower(w)
		wordFrequencyMap[lowered]++
		if wordFrequencyMap[lowered] > maxCount {
			maxCount = wordFrequencyMap[lowered]
		}
	}
	var wordCountsArr = make([]wordCount, 0)
	for k, v := range wordFrequencyMap {
		wordCountsArr = append(wordCountsArr, wordCount{
			word:  k,
			count: v,
		})
	}
	sort.Slice(wordCountsArr, func(i, j int) bool {
		return wordCountsArr[i].count > wordCountsArr[j].count
	})
	for _, item := range wordCountsArr {
		res = append(res, item.word)
		if len(res) >= max {
			break
		}
	}
	return res
}

//MostFrequentWithCustomQuickSort return 'max' most frequent words in 'text' - same algo as MostFrequentWithSliceSort but adjusted quicksortDesc func
func MostFrequentWithCustomQuickSort(text string, max int) []string {
	var res = make([]string, 0)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	var wordFrequencyMap = make(map[string]int, 0) // map["word1": 2, "word2": 1]
	var maxCount int
	for _, w := range words {
		var lowered = strings.ToLower(w)
		wordFrequencyMap[lowered]++
		if wordFrequencyMap[lowered] > maxCount {
			maxCount = wordFrequencyMap[lowered]
		}
	}
	var wordCountsArr = make([]wordCount, 0)
	for k, v := range wordFrequencyMap {
		wordCountsArr = append(wordCountsArr, wordCount{
			word:  k,
			count: v,
		})
	}
	quicksortDesc(wordCountsArr) // replaces the sort.Slice() call (get rid of reflection)
	for _, item := range wordCountsArr {
		res = append(res, item.word)
		if len(res) >= max {
			break
		}
	}
	return res
}

func quicksortDesc(a []wordCount) []wordCount {
	if len(a) < 2 {
		return a
	}

	if len(a) < 12 {
		var n = len(a)
		for i := 1; i < n; i++ {
			j := i
			for j > 0 {
				if a[j-1].count < a[j].count {
					a[j-1], a[j] = a[j], a[j-1]
				}
				j = j - 1
			}
		}
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Intn(len(a))

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].count > a[right].count {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksortDesc(a[:left])
	quicksortDesc(a[left+1:])

	return a
}
