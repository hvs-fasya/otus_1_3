package otus_1_3

import (
	"testing"
)

func BenchmarkFrequncyCounter(b *testing.B) {
	text := `a: a, B, b, c, c, d, d, e, e, f, g!... `
	var n = 5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MostFrequent(text, n)
	}
}

func BenchmarkFrequncyCounterWithSliceSort(b *testing.B) {
	text := `a: a, B, b, c, c, d, d, e, e, f, g!... `
	var n = 5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MostFrequentWithSliceSort(text, n)
	}
}

func BenchmarkFrequncyCounterWithCustomQuickSort(b *testing.B) {
	text := `a: a, B, b, c, c, d, d, e, e, f, g!... `
	var n = 5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MostFrequentWithCustomQuickSort(text, n)
	}
}
