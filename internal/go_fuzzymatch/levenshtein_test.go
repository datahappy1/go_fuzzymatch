package goFuzzymatch

import (
	"strings"
	"testing"
)

func TestLevenshteinRatio(t *testing.T) {
	ans := LevenshteinRatio(strings.ToLower("Apple Inc."), strings.ToLower("apple Inc"))
	if ans != 94 {
		t.Errorf("levenshteinRatio(\"Apple Inc.\", \"apple Inc\") = %d; want 94", ans)
	}
}
