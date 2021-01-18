package fm

import (
	match "github.com/datahappy1/go_fuzzymatch/internal"
)

// FuzzyMatch returns uint16
func FuzzyMatch(string1 string, string2 string, mode string) uint16 {

	if string1 == string2 {
		return 1
	} else if string1 == "" || string2 == "" {
		return 0
	} else {
		var m = &match.Match{}
		if mode == "simple" {
			m.Strategy = match.Simple{}
			return m.MatchStrings(string1, string2)
		} else if mode == "deepDive" {
			m.Strategy = match.DeepDive{}
			return m.MatchStrings(string1, string2)
		} else if mode == "combined" {
			m.Strategy = match.Combined{}
			return m.MatchStrings(string1, string2)
		} else {
			return 0
		}
	}
}
