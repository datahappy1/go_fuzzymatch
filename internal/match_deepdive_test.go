package match

import (
	"fmt"
	"testing"
)

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
