package pkg

import (
	"github.com/datahappy1/go_fuzzymatch/internal/go_fuzzymatch"
)

// FuzzyMatch returns int
func FuzzyMatch(string1 string, string2 string, mode string) int {

	if string1 == string2 {
		return 100
	} else if string1 == "" || string2 == "" {
		return 0
	} else {
		var m = &go_fuzzymatch.Match{}
		if mode == "simple" {
			m.Strategy = go_fuzzymatch.Simple{}
			return m.MatchStrings(string1, string2)
		} else if mode == "deepDive" {
			m.Strategy = go_fuzzymatch.DeepDive{}
			return m.MatchStrings(string1, string2)
		} else if mode == "combined" {
			m.Strategy = go_fuzzymatch.Combined{}
			return m.MatchStrings(string1, string2)
		} else {
			return -1
		}
	}
}
