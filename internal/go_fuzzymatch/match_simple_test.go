package gofuzzymatch

import (
	"fmt"
	"testing"
)

func TestCalculateLevenshteinForIterations(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		want   []int
	}{
		{s1: "aplle", s2: "tree", want: []int{22}},
		{s1: "apple", s2: "apple inc", want: []int{100, 0}},
		{s1: "apple", s2: "Apple Inc.", want: []int{100, 0}},
		{s1: "Apple", s2: "Apple Inc", want: []int{100, 0}},
		{s1: "aplle", s2: "Apple", want: []int{80}},
		{s1: "Apple", s2: "Apple Corp. GMBH", want: []int{100, 22, 0}},
		{s1: "Apple", s2: "GMBH Apple Corp", want: []int{0, 100, 22}},
		{s1: "Aplle", s2: "GMBH Apple Corp", want: []int{0, 80, 22}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.s1, tt.s2)

		staticString := createEvaluatedString(tt.s1)
		permutableString := createEvaluatedString(tt.s2)

		t.Run(testname, func(t *testing.T) {
			ans := calculateLevenshteinForIterations(*staticString, *permutableString)
			if len(ans) != len(tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
				t.FailNow()
			}
			for i, v := range ans {
				if v != tt.want[i] {
					t.Errorf("got %d, want %d", ans, tt.want)
				}
			}
		})
	}
}

func TestMatchSimple(t *testing.T) {
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
		{s1: "Apple Corp.", s2: "Apple Corp. GMBH", want: 80},
		{s1: "GMBH Apple Corp", s2: "Apple Inc.", want: 58},
		{s1: "apple Inc.", s2: "GMBH Apple Corp.", want: 58},
		{s1: "aplle Inc.", s2: "GMBH Apple Corp.", want: 50},
	}
	var m = &Match{}
	m.Strategy = Simple{}

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
