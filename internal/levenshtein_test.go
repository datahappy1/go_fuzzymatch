package match

import (
	"strings"
	"testing"
)

func TestLevenshteinRatio(t *testing.T) {
	ans := LevenshteinRatio(strings.ToLower("Apple Inc."), strings.ToLower("apple Inc"))
	if ans != 0.94736844 {
		t.Errorf("levenshteinRatio(\"Apple Inc.\", \"apple Inc\") = %g; want 0.94736844", ans)
	}
}
