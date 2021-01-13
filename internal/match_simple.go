package match

import (
	"sort"
	"strings"
)

// Simple returns struct
type Simple struct{}

func calculateLevenshteinForIterations(staticString evaluatedString, permutableString evaluatedString) []float32 {
	var outputSlice []float32

	for _, splitStaticStringItem := range staticString.valueByWordSplitArray {
		outputSlice = append(outputSlice, LevenshteinRatio(splitStaticStringItem, permutableString.value))
	}

	return outputSlice
}

func (Simple) matchStrings(s1 string, s2 string) float32 {
	String1 := createEvaluatedString(s1)
	String2 := createEvaluatedString(s2)
	var outputSlice []float32

	if String1.valueByWordSplitArrayLength > 1 && String2.valueByWordSplitArrayLength > 1 {
		sort.Strings(String1.valueByWordSplitArray)
		sort.Strings(String2.valueByWordSplitArray)
		return LevenshteinRatio(strings.Join(String1.valueByWordSplitArray, " "),
			strings.Join(String2.valueByWordSplitArray, " "))
	} else if String1.valueByWordSplitArrayLength > 1 && String2.valueByWordSplitArrayLength == 1 {
		outputSlice = calculateLevenshteinForIterations(*String1, *String2)
		return maxOfSliceOfFloats(outputSlice)
	} else if String1.valueByWordSplitArrayLength == 1 && String2.valueByWordSplitArrayLength > 1 {
		outputSlice = calculateLevenshteinForIterations(*String2, *String1)
		return maxOfSliceOfFloats(outputSlice)
	} else {
		return LevenshteinRatio(String1.value, String2.value)
	}
}
