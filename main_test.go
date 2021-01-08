package main

import (
	"strings"
	"testing"
)

func TestLevenshteinRatioAndDistance(t *testing.T) {
	ans := LevenshteinRatioAndDistance(strings.ToLower("Apple Inc."), strings.ToLower("apple Inc"), true)
	if ans != "0.9473684210526315" {
		t.Errorf("levenshteinRatioAndDistance(strings.ToLower(\"Apple Inc.\"), strings.ToLower(\"apple Inc\"), true) = %s; want \"0.9473684210526315\"", ans)
	}
}
