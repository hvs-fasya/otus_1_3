package otus_1_3

import (
	"testing"
)

func TestMostFrequent(t *testing.T) {
	var n = 5
	tests := []struct {
		text string
		want []string
	}{
		{`a a B b c c d d e e f g `, []string{`a`, `b`, `c`, `d`, `e`}},
		{`a a B b c c d d e e  f f g g  h h i i i j j j k `, []string{`i`, `j`, `a`, `b`, `c`, `d`, `e`, `f`, `g`, `h`}},
		{``, []string{}},
	}
	for _, test := range tests {
		res := MostFrequent(test.text, n)
		if !match(res, test.want) {
			t.Errorf("MostFrequent() = %q, want %q", res, test.want)
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
