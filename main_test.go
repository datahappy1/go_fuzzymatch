package main_test

import (
	"strings"
	"testing"
)

func testLevenshteinRatioAndDistance(t *testing.T) {
	ans := LevenshteinRatioAndDistance(strings.ToLower("Apple Inc."), strings.ToLower("apple Inc"), true)
	if ans != 0.9473684210526315 {
		t.Errorf("levenshteinRatioAndDistance(strings.ToLower('Apple Inc.'), strings.ToLower('apple Inc'), true) = %d; want 0.9473684210526315", ans)
	}
}
