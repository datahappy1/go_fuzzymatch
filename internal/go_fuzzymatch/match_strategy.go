package go_fuzzymatch

type strategy interface {
	matchStrings(s1 string, s2 string) int
}
