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
		{s1: "Apple", s2: "Apple Inc", want: []float32{1, 0}},
		{s1: "aplle", s2: "Apple", want: []float32{0.8}},
		{s1: "Apple Corp.", s2: "Apple Corp. GMBH", want: []float32{1, 0.5, 0.5, 0.6, 0.42105263, 0.5263158}},
		{s1: "Apple Inc.", s2: "GMBH Apple Corp", want: []float32{0.5263158, 0.6315789, 0.22222222, 0.22222222, 0.7368421, 0.5263158}},
		{s1: "Aplle Inc.", s2: "GMBH Apple Corp", want: []float32{0.42105263, 0.5263158, 0.22222222, 0.22222222, 0.6315789, 0.42105263}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.s1, tt.s2)

		staticString := createEvaluatedString(tt.s1)
		permutableString := createEvaluatedString(tt.s2)

		t.Run(testname, func(t *testing.T) {
			ans := calculateLevenshteinForPermutations(*staticString, *permutableString)
			if len(ans) != len(tt.want) {
				t.Errorf("got %g, want %g", ans, tt.want)
				t.FailNow()
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
		{s1: "apple inc", s2: "apple inc", want: 1},
		{s1: "apple inc", s2: "Apple Inc.", want: 1},
		{s1: "Apple", s2: "Apple Inc.", want: 1},
		{s1: "Apple Inc", s2: "Apple", want: 1},
		{s1: "aplle", s2: "Apple", want: 0.8},
		{s1: "Apple Corp.", s2: "Apple Corp. GMBH", want: 1},
		{s1: "GMBH Apple Corp", s2: "Apple Inc.", want: 0.7368421},
		{s1: "apple Inc.", s2: "GMBH Apple Corp.", want: 0.7368421},
		{s1: "aplle Inc.", s2: "GMBH Apple Corp.", want: 0.6315789},
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
