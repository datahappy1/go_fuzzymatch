package gofuzzymatch

import (
	"fmt"
	"testing"
)

func TestMatchCombined(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		want   int
	}{
		{s1: "aplle", s2: "tree", want: 22},
		{s1: "apple inc", s2: "apple inc", want: 100},
		{s1: "apple inc", s2: "Apple Inc.", want: 100},
		{s1: "Apple", s2: "Apple Inc.", want: 100},
		{s1: "Apple Inc", s2: "Apple", want: 100},
		{s1: "aplle", s2: "Apple", want: 80},
		{s1: "Apple Corp.", s2: "Apple Corp. GMBH", want: 100},
		{s1: "GMBH Apple Corp", s2: "Apple Inc.", want: 73},
		{s1: "apple Inc.", s2: "GMBH Apple Corp.", want: 73},
		{s1: "aplle Inc.", s2: "GMBH Apple Corp.", want: 63},
	}
	var m = &Match{}
	m.Strategy = Combined{}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.s1, tt.s2)
		t.Run(testname, func(t *testing.T) {
			ans := m.MatchStrings(tt.s1, tt.s2)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
