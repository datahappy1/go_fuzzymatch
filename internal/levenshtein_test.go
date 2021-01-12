package match

import (
	"strings"
	"testing"
)

func TestLevenshteinRatio(t *testing.T) {
	ans := LevenshteinRatio(strings.ToLower("Apple Inc."), strings.ToLower("apple Inc"))
	if ans != 0.9473684210526315 {
		t.Errorf("levenshteinRatio(\"Apple Inc.\", \"apple Inc\") = %g; want 0.9473684210526315", ans)
	}
}
