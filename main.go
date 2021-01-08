package main

import (
	"flag"
	"fmt"
	"strings"
)

func convertToFloat(i int) float64 {
	return float64(i)
}

func minOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func splitStringToArrayByWhitespace(s string) []string {
	var splitStringArray []string
	for _, word := range strings.Fields(s) {
		splitStringArray = append(splitStringArray, word)
	}
	return splitStringArray
}

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func getPermutations() []string {
	d := permutations("ABCD")
	fmt.Print(d)
	return d
}

//func RemoveNonAlphaChars(s string) string{
//	return s
//}

// Levenshtein Ratio and Distance function
func LevenshteinRatioAndDistance(s1 string, s2 string, ratioCalc bool) string {
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
	// fmt.Println(distance)

	// Populate matrix of zeros with the indeces of each character of both strings
	for i := 1; i < rows; i++ {
		for ii := 1; ii < cols; ii++ {
			distance[i][0] = i
			distance[0][ii] = ii
		}
	}
	//fmt.Println(distance)

	// Iterate over the matrix to compute the cost of deletions,insertions and/or substitutions
	for col := 1; col < cols; col++ {
		//fmt.Printf("c %v\n", col)

		for row := 1; row < rows; row++ {
			//fmt.Printf("r %v\n", row)

			if s1[row-1] == s2[col-1] {
				cost = 0 // If the characters are the same in the two strings in a given position [i,j] then the cost is 0
			} else {
				// In order to align the results with those of the Python Levenshtein package, if we choose to calculate the ratio
				// the cost of a substitution is 2. If we calculate just distance, then the cost of a substitution is 1.
				if ratioCalc == true {
					cost = 2
				} else {
					cost = 1
				}
			}

			distance[row][col] = minOf(
				distance[row-1][col]+1,      // Cost of deletions
				distance[row][col-1]+1,      // Cost of insertions
				distance[row-1][col-1]+cost) // Cost of substitutions

		}
	}

	if ratioCalc == true {
		// Computation of the Levenshtein Distance Ratio
		ratio := (convertToFloat(rowLength+colLength) - convertToFloat(distance[rowLength][colLength])) /
			convertToFloat(rowLength+colLength)
		return fmt.Sprintf("%g", ratio)
	}
	// fmt.Println(distance)
	// Uncomment if you want to see the matrix showing how the algorithm computes the cost of deletions,
	// insertions and/or substitutions
	// This is the minimum number of edits needed to convert string a to string b
	return fmt.Sprintf("The strings are %d edits away", distance[rowLength][colLength])

}

func main() {
	string1Ptr := flag.String("string1", "a", "first string")
	string2Ptr := flag.String("string2", "b", "second string")
	ratioCalcPtr := flag.Bool("ratioCalc", true, "calculate ratio")
	flag.Parse()

	//xs1 := RemoveNonAlphaChars(*string1Ptr)
	//xs2 := RemoveNonAlphaChars(*string2Ptr)
	// continue working with strings without non alpha chars

	fmt.Println(splitStringToArrayByWhitespace(*string1Ptr))
	if len(splitStringToArrayByWhitespace(*string1Ptr)) > 1 && len(splitStringToArrayByWhitespace(*string2Ptr)) > 1 {
		string1permutations := getPermutations(splitStringToArrayByWhitespace(*string1Ptr))
		string2permutations := getPermutations(splitStringToArrayByWhitespace(*string2Ptr))
	}
	fmt.Println(LevenshteinRatioAndDistance(strings.ToLower(*string1Ptr), strings.ToLower(*string2Ptr), *ratioCalcPtr))
}
