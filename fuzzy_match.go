package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	prmt "github.com/datahappy1/permutation"
)

type evaluatedString struct {
	value                       string
	valueByWordSplitArray       []string
	valueByWordSplitArrayLength int
}

func createEvaluatedString(v string) *evaluatedString {
	processedInputString := strings.ToLower(removeUnusedChars(v))
	stringWordSplit := splitStringToUniqueValuesSliceByWhitespace(processedInputString)
	s := evaluatedString{
		value:                       processedInputString,
		valueByWordSplitArray:       stringWordSplit,
		valueByWordSplitArrayLength: len(stringWordSplit)}
	return &s
}

func checkInputStringsEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return false
}

func convertIntToFloat(i int) float32 {
	return float32(i)
}

func minOfVarsOfInts(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func maxOfSliceOfFloats(slice []float32) float32 {
	max := slice[0]
	for _, i := range slice {
		if max < i {
			max = i
		}
	}
	return max
}

func appendStringToSliceIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func splitStringToUniqueValuesSliceByWhitespace(s string) []string {
	var splitStringArray []string
	for _, word := range strings.Fields(s) {
		splitStringArray = appendStringToSliceIfMissing(splitStringArray, word)
	}
	return splitStringArray
}

func removeUnusedChars(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
}

// LevenshteinRatio returns float32
func LevenshteinRatio(s1 string, s2 string) float32 {
	var rowLength int = len(s1)
	var colLength int = len(s2)
	var rows int = rowLength + 1
	var cols int = colLength + 1
	var cost int = 0

	// Initialize a rows length slice of empty slices
	distance := make([][]int, rows)

	// Initialize the cols empty slices
	for x := 0; x < rows; x++ {
		distance[x] = make([]int, cols)
	}

	// Populate matrix of zeros with the indeces of each character of both strings
	for i := 1; i < rows; i++ {
		for ii := 1; ii < cols; ii++ {
			distance[i][0] = i
			distance[0][ii] = ii
		}
	}

	// Iterate over the matrix to compute the cost of deletions,insertions and/or substitutions
	for col := 1; col < cols; col++ {
		for row := 1; row < rows; row++ {
			if s1[row-1] == s2[col-1] {
				cost = 0 // If the characters are the same in the two strings in a given position [i,j] then the cost is 0
			} else {
				// In order to align the results with those of the Python Levenshtein package, if we choose to calculate the ratio
				// the cost of a substitution is 2.
				cost = 2
			}

			distance[row][col] = minOfVarsOfInts(
				distance[row-1][col]+1,      // Cost of deletions
				distance[row][col-1]+1,      // Cost of insertions
				distance[row-1][col-1]+cost) // Cost of substitutions

		}
	}

	// Computation of the Levenshtein Distance Ratio
	ratio := (convertIntToFloat(rowLength+colLength) - convertIntToFloat(distance[rowLength][colLength])) /
		convertIntToFloat(rowLength+colLength)

	return ratio
}

// Match returns float32
func Match(s1 string, s2 string) float32 {
	if checkInputStringsEqual(s1, s2) == true {
		return 1
	}

	String1 := createEvaluatedString(s1)
	String2 := createEvaluatedString(s2)
	outputSlice := []float32{}

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

// MatchDeepDive returns float32
func MatchDeepDive(s1 string, s2 string) float32 {
	if checkInputStringsEqual(s1, s2) == true {
		return 1
	}

	String1 := createEvaluatedString(s1)
	String2 := createEvaluatedString(s2)
	outputSlice := []float32{}

	if String1.valueByWordSplitArrayLength < String2.valueByWordSplitArrayLength {
		splitString2ArraySliced := String2.valueByWordSplitArray[0:String1.valueByWordSplitArrayLength]
		splitString2ArraySlicedJoined := strings.Join(splitString2ArraySliced, " ")
		p := prmt.New(prmt.StringSlice(String1.valueByWordSplitArray))
		for p.Next() {
			outputSlice = append(outputSlice, LevenshteinRatio(strings.Join(String1.valueByWordSplitArray, " "), splitString2ArraySlicedJoined))
		}
	} else {
		splitString1ArraySliced := String1.valueByWordSplitArray[0:String2.valueByWordSplitArrayLength]
		splitString1ArraySlicedJoined := strings.Join(splitString1ArraySliced, " ")
		p := prmt.New(prmt.StringSlice(String2.valueByWordSplitArray))
		for p.Next() {
			outputSlice = append(outputSlice, LevenshteinRatio(strings.Join(String2.valueByWordSplitArray, " "), splitString1ArraySlicedJoined))
		}
	}
	return maxOfSliceOfFloats(outputSlice)
}

func main() {
	string1Ptr := flag.String("string1", "apple", "first string")
	string2Ptr := flag.String("string2", "bear", "second string")
	deepDive := flag.Bool("deepDive", false, "deep dive")
	flag.Parse()

	if *deepDive == true {
		fmt.Println(MatchDeepDive(*string1Ptr, *string2Ptr))
	} else {
		fmt.Println(Match(*string1Ptr, *string2Ptr))
	}
}
