package match

import (
	"strings"

	prmt "github.com/datahappy1/permutation"
)

// DeepDive returns struct
type DeepDive struct{}

func calculateLevenshteinForPermutations(staticString evaluatedString, permutableString evaluatedString) []int {
	var outputSlice []int
	var i = 0
	var maxIterSize = 68000

	p := prmt.New(prmt.StringSlice(permutableString.valueByWordSplitArray))
	for p.Next() {
		outputSlice = append(outputSlice, LevenshteinRatio(strings.Join(staticString.valueByWordSplitArray, " "),
			strings.Join(permutableString.valueByWordSplitArray[0:staticString.valueByWordSplitArrayLength], " ")))
		i++
		if i > maxIterSize {
			break
		}
	}
	return outputSlice
}

func (DeepDive) matchStrings(s1 string, s2 string) int {
	String1 := createEvaluatedString(s1)
	String2 := createEvaluatedString(s2)
	var outputSlice []int

	if String1.valueByWordSplitArrayLength == 1 && String2.valueByWordSplitArrayLength == 1 {
		return LevenshteinRatio(String1.value, String2.value)
	} else if String1.valueByWordSplitArrayLength < String2.valueByWordSplitArrayLength {
		outputSlice = calculateLevenshteinForPermutations(*String1, *String2)
	} else {
		outputSlice = calculateLevenshteinForPermutations(*String2, *String1)
	}
	return maxOfSliceOfInts(outputSlice)
}
