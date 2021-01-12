package match

import (
	"sort"
	"strings"
)

// Simple returns struct
type Simple struct{}

func (Simple) matchStrings(s1 string, s2 string) float32 {
	if checkInputStringsEqual(s1, s2) == true {
		return 1
	}

	String1 := createEvaluatedString(s1)
	String2 := createEvaluatedString(s2)
	var outputSlice []float32

	if String1.valueByWordSplitArrayLength > 1 && String2.valueByWordSplitArrayLength > 1 {
		sort.Strings(String1.valueByWordSplitArray)
		sort.Strings(String2.valueByWordSplitArray)
		return LevenshteinRatio(strings.Join(String1.valueByWordSplitArray, " "), strings.Join(String2.valueByWordSplitArray, " "))
	} else if String1.valueByWordSplitArrayLength > 1 && String2.valueByWordSplitArrayLength == 1 {
		for _, splitString1Item := range String1.valueByWordSplitArray {
			outputSlice = append(outputSlice, LevenshteinRatio(splitString1Item, String2.value))
		}
		return maxOfSliceOfFloats(outputSlice)
	} else if String1.valueByWordSplitArrayLength == 1 && String2.valueByWordSplitArrayLength > 1 {
		for _, splitString2Item := range String2.valueByWordSplitArray {
			outputSlice = append(outputSlice, LevenshteinRatio(String1.value, splitString2Item))
		}
		return maxOfSliceOfFloats(outputSlice)
	} else {
		return LevenshteinRatio(String1.value, String2.value)
	}
}
