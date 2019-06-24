package otus_1_3

import (
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestMostFrequent(t *testing.T) {
	var n = 5
	tests := []struct {
		text string
		want []string
	}{
		{`a a B b c c d d e e f g `, []string{`a`, `b`, `c`, `d`, `e`}},
		{`a: a, B, b, c, c, d, d, e, e, f, g!... `, []string{`a`, `b`, `c`, `d`, `e`}},
		{`a a B b c c d d e e  f f g g  h h i i i j j j k `, []string{`i`, `j`, `a`, `b`, `c`, `d`, `e`, `f`, `g`, `h`}},
		{``, []string{}},
	}
	for _, test := range tests {
		res := MostFrequent(test.text, n)
		if !match(res, test.want) {
			t.Errorf("MostFrequent() = %q, want %q", res, test.want)
		}
		res1 := MostFrequentWithSliceSort(test.text, n)
		if !match(res1, test.want) {
			t.Errorf("MostFrequentWithSliceSort() = %q, want %q", res1, test.want)
		}
		res2 := MostFrequentWithCustomQuickSort(test.text, n)
		if !match(res2, test.want) {
			t.Errorf("MostFrequentWithCustomQuickSort() = %q, want %q", res2, test.want)
		}
	}
}

//result and want should not be equal - elements of result should be members if the want set
func match(result []string, want []string) bool {
	for _, res := range result {
		var found bool
		for _, w := range want {
			if w == res {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func Test_quicksort(t *testing.T) {
	type args struct {
		a []wordCount
	}
	//generate many wordCount items slice with unique counts - unique to test assertion
	var manyItems []wordCount
	var countsUsed = make([]int, 0)
	for i := 0; i < 25; i++ {
		var count int
		var used = true
		for used {
			count = rand.Intn(100)
			var found bool
			for _, item := range countsUsed {
				if item == count {
					found = true
					break
				}
			}
			if found {
				continue
			}
			used = false
		}
		countsUsed = append(countsUsed, count)
		manyItems = append(manyItems, wordCount{"word" + strconv.Itoa(i), count})
	}
	var manyItemsSorted = append([]wordCount{}, manyItems...) // assert to sorted by sort package func - we trust sort package!!!!
	sort.Slice(manyItemsSorted, func(i, j int) bool {
		return manyItemsSorted[i].count > manyItemsSorted[j].count
	})
	tests := []struct {
		name string
		args args
		want []wordCount
	}{
		{"empty slice", args{make([]wordCount, 0)}, []wordCount{}},
		{
			"few elements in slice",
			args{[]wordCount{{"word1", 15}, {"word2", 1}, {"word3", 5}}},
			[]wordCount{{"word1", 15}, {"word3", 5}, {"word2", 1}},
		},
		{"many elements in slice", args{manyItems}, manyItemsSorted},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quicksortDesc(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quicksortDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}
