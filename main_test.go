package main

import (
	"strings"
	"testing"
)

func TestLevenshteinRatio(t *testing.T) {
	ans := LevenshteinRatio(strings.ToLower("Apple Inc."), strings.ToLower("apple Inc"))
	if ans != 0.9473684210526315 {
		t.Errorf("levenshteinRatio(strings.ToLower(\"Apple Inc.\"), strings.ToLower(\"apple Inc\"), true) = %g; want 0.9473684210526315", ans)
	}
}
