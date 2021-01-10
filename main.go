package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	prmt "github.com/datahappy1/permutation"
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

func maxOf(slice []float64) float64 {
	max := slice[0]
	for _, i := range slice {
		if max < i {
			max = i
		}
	}
	return max
}

func splitStringToArrayByWhitespace(s string) []string {
	var splitStringArray []string
	for _, word := range strings.Fields(s) {
		splitStringArray = append(splitStringArray, word)
	}
	return splitStringArray
}

// RemoveNonAlphaChars returns string
func RemoveNonAlphaChars(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
}

// LevenshteinRatio returns string
func LevenshteinRatio(s1 string, s2 string) float64 {
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
				// the cost of a substitution is 2.
				cost = 2
			}

			distance[row][col] = minOf(
				distance[row-1][col]+1,      // Cost of deletions
				distance[row][col-1]+1,      // Cost of insertions
				distance[row-1][col-1]+cost) // Cost of substitutions

		}
	}

	// Computation of the Levenshtein Distance Ratio
	ratio := (convertToFloat(rowLength+colLength) - convertToFloat(distance[rowLength][colLength])) /
		convertToFloat(rowLength+colLength)
	//return fmt.Sprintf("%g", ratio)
	return ratio
}

func main() {
	string1Ptr := flag.String("string1", "a", "first string")
	string2Ptr := flag.String("string2", "b", "second string")
	flag.Parse()

	String1 := strings.ToLower(RemoveNonAlphaChars(*string1Ptr))
	String2 := strings.ToLower(RemoveNonAlphaChars(*string2Ptr))

	//fmt.Println(String1)
	//fmt.Println(splitStringToArrayByWhitespace(String1))

	// start permutation stuff
	splitString1 := splitStringToArrayByWhitespace(String1)
	splitString2 := splitStringToArrayByWhitespace(String2)

	outputSlice := []float64{}

	if len(splitString1) > 1 && len(splitString2) > 1 {
		p := prmt.New(prmt.StringSlice(splitString1))
		for p.Next() {
			// fmt.Println(splitString1)
			// fmt.Println(strings.Join(splitString1, " "))
			// fmt.Println(String2)
			// fmt.Println(LevenshteinRatioAndDistance(strings.Join(splitString1, " "), String2, *ratioCalcPtr))
			outputSlice = append(outputSlice, LevenshteinRatio(strings.Join(splitString1, " "), String2))

			// fmt.Println("---")
		}
		fmt.Println(maxOf(outputSlice))
	} else {
		fmt.Println(LevenshteinRatio(String1, String2))
	}
}
