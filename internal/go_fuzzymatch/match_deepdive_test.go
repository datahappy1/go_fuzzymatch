package gofuzzymatch

import (
	"fmt"
	"testing"
)

func TestCalculateLevenshteinForPermutations(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		want   []int
	}{
		{s1: "aplle", s2: "tree", want: []int{22}},
		{s1: "apple inc", s2: "apple inc", want: []int{100, 55}},
		{s1: "apple inc", s2: "Apple Inc.", want: []int{100, 55}},
		{s1: "Apple", s2: "Apple Inc", want: []int{100, 0}},
		{s1: "aplle", s2: "Apple", want: []int{80}},
		{s1: "Apple Corp.", s2: "Apple Corp. GMBH", want: []int{100, 50, 50, 60, 42, 52}},
		{s1: "Apple Inc.", s2: "GMBH Apple Corp", want: []int{52, 63, 22, 22, 73, 52}},
		{s1: "Aplle Inc.", s2: "GMBH Apple Corp", want: []int{42, 52, 22, 22, 63, 42}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.s1, tt.s2)

		staticString := createEvaluatedString(tt.s1)
		permutableString := createEvaluatedString(tt.s2)

		t.Run(testname, func(t *testing.T) {
			ans := calculateLevenshteinForPermutations(*staticString, *permutableString)
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

func TestMatchDeepDive(t *testing.T) {
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
	m.Strategy = DeepDive{}

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
