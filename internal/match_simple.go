package match

import (
	"sort"
	"strings"
)

// Simple returns struct
type Simple struct{}

func calculateLevenshteinForIterations(staticString evaluatedString, iterableString evaluatedString) []uint16 {
	var outputSlice []uint16

	for _, iterableStringItem := range iterableString.valueByWordSplitSlice {
		outputSlice = append(outputSlice, LevenshteinRatio(iterableStringItem, staticString.value))
	}

	return outputSlice
}

func (Simple) matchStrings(s1 string, s2 string) uint16 {
	String1 := createEvaluatedString(s1)
	String2 := createEvaluatedString(s2)
	var outputSlice []uint16

	if String1.valueByWordSplitSliceLength > 1 && String2.valueByWordSplitSliceLength > 1 {
		sort.Strings(String1.valueByWordSplitSlice)
		sort.Strings(String2.valueByWordSplitSlice)
		return LevenshteinRatio(strings.Join(String1.valueByWordSplitSlice, " "),
			strings.Join(String2.valueByWordSplitSlice, " "))
	} else if String1.valueByWordSplitSliceLength == 1 && String2.valueByWordSplitSliceLength > 1 {
		outputSlice = calculateLevenshteinForIterations(*String1, *String2)
		return maxOfSliceOfUnsignedInts(outputSlice)
	} else if String1.valueByWordSplitSliceLength > 1 && String2.valueByWordSplitSliceLength == 1 {
		outputSlice = calculateLevenshteinForIterations(*String2, *String1)
		return maxOfSliceOfUnsignedInts(outputSlice)
	} else {
		return LevenshteinRatio(String1.value, String2.value)
	}
}
