package match

import (
	"strings"

	prmt "github.com/datahappy1/permutation"
)

// DeepDive returns struct
type DeepDive struct{}

func (DeepDive) matchStrings(s1 string, s2 string) float32 {
	if checkInputStringsEqual(s1, s2) == true {
		return 1
	}

	String1 := createEvaluatedString(s1)
	String2 := createEvaluatedString(s2)
	var outputSlice []float32
	var i = 0
	var maxIterSize = 70000

	if String1.valueByWordSplitArrayLength < String2.valueByWordSplitArrayLength {
		p := prmt.New(prmt.StringSlice(String2.valueByWordSplitArray))
		for p.Next() {
			outputSlice = append(outputSlice, LevenshteinRatio(strings.Join(String1.valueByWordSplitArray, " "),
				strings.Join(String2.valueByWordSplitArray[0:String1.valueByWordSplitArrayLength], " ")))
			i++
			if i > maxIterSize {
				break
			}
		}
	} else {
		p := prmt.New(prmt.StringSlice(String1.valueByWordSplitArray))
		for p.Next() {
			outputSlice = append(outputSlice, LevenshteinRatio(strings.Join(String2.valueByWordSplitArray, " "),
				strings.Join(String1.valueByWordSplitArray[0:String2.valueByWordSplitArrayLength], " ")))
			i++
			if i > maxIterSize {
				break
			}
		}
	}
	return maxOfSliceOfFloats(outputSlice)
}
