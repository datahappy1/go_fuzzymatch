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
		{"apple inc", "Apple Inc.", 1.0},
		{"aplle", "Apple Corp.", 0.5},
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
