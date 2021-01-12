package match

import (
	"fmt"
	"testing"
)

func TestCalculateLevenshteinForPermutations(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		want   []float32
	}{
		{s1: "apple inc", s2: "apple inc", want: []float32{1, 0.5555556}},
		{s1: "apple inc", s2: "Apple Inc.", want: []float32{1, 0.5555556}},
		{s1: "Apple", s2: "Apple Inc.", want: []float32{1, 0}},
		{s1: "Apple Inc", s2: "Apple", want: []float32{1}},
		{s1: "aplle", s2: "Apple", want: []float32{1}},
		{s1: "Apple Corp.", s2: "Apple Corp. GMBH", want: []float32{1, 2, 3}},
		{s1: "GMBH Apple Corp", s2: "Apple Inc.", want: []float32{1, 2, 3}},
		{s1: "apple Inc.", s2: "GMBH Apple Corp.", want: []float32{1, 2, 3}},
		{s1: "aplle Inc.", s2: "GMBH Apple Corp.", want: []float32{1, 2, 3}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.s1, tt.s2)

		staticString := createEvaluatedString(tt.s1)
		permutableString := createEvaluatedString(tt.s2)

		t.Run(testname, func(t *testing.T) {
			ans := calculateLevenshteinForPermutations(*staticString, *permutableString)
			if len(ans) != len(tt.want) {
				t.Errorf("got %g, want %g", ans, tt.want)
			}
			for i, v := range ans {
				if v != tt.want[i] {
					t.Errorf("got %g, want %g", ans, tt.want)
				}
			}
		})
	}
}

func TestMatchDeepDive(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		want   float32
	}{
		{"apple inc", "apple inc", 1},
		{"apple inc", "Apple Inc.", 1},
		{"Apple", "Apple Inc.", 1},
		{"Apple Inc", "Apple", 1},
		{"aplle", "Apple", 0.8},
		{"Apple Corp.", "Apple Corp. GMBH", 1},
		{"GMBH Apple Corp", "Apple Inc.", 0.7368421},
		{"apple Inc.", "GMBH Apple Corp.", 0.7368421},
		{"aplle Inc.", "GMBH Apple Corp.", 0.6315789},
	}
	var m = &Match{}
	m.Strategy = DeepDive{}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.s1, tt.s2)
		t.Run(testname, func(t *testing.T) {
			ans := m.MatchStrings(tt.s1, tt.s2)
			if ans != tt.want {
				t.Errorf("got %g, want %g", ans, tt.want)
			}
		})
	}
}
