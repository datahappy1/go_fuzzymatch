package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestLevenshteinRatio(t *testing.T) {
	ans := LevenshteinRatio(strings.ToLower("Apple Inc."), strings.ToLower("apple Inc"))
	if ans != 0.9473684210526315 {
		t.Errorf("levenshteinRatio(strings.ToLower(\"Apple Inc.\"), strings.ToLower(\"apple Inc\"), true) = %g; want 0.9473684210526315", ans)
	}
}

func TestMatch(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		want   float32
	}{
		{"apple inc", "apple inc", 1},
		{"apple inc", "Apple Inc.", 1},
		{"Apple", "Apple Inc.", 1},
		{"Apple Inc", "Apple", 1},
		{"aplle", "Apple", 0.8},
		{"Apple Corp.", "Apple Corp. GMBH", 0.8},
		{"GMBH Apple Corp", "Apple Inc.", 0.5833333},
		{"apple Inc.", "GMBH Apple Corp.", 0.5833333},
		{"aplle Inc.", "GMBH Apple Corp.", 0.5},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.s1, tt.s2)
		t.Run(testname, func(t *testing.T) {
			ans := Match(tt.s1, tt.s2)
			if ans != tt.want {
				t.Errorf("got %g, want %g", ans, tt.want)
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
		{"GMBH Apple Corp", "Apple Inc.", 0.6315789},
		{"apple Inc.", "GMBH Apple Corp.", 0.6315789},
		{"aplle Inc.", "GMBH Apple Corp.", 0.5263158},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.s1, tt.s2)
		t.Run(testname, func(t *testing.T) {
			ans := MatchDeepDive(tt.s1, tt.s2)
			if ans != tt.want {
				t.Errorf("got %g, want %g", ans, tt.want)
			}
		})
	}
}
