package main

import (
	"fmt"
	"strconv"
)

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func levenshtein_ratio_and_distance(s string, t string, ratio_calc bool) string {
	// Initialize helper variables
	var cost int = 0
	var row int = len(s)
	var col int = len(t)

	// Initialize matrix of zeros
	var rows int = len(s) + 1
	var cols int = len(t) + 1

	// Initialize a ten length slice of empty slices
	distance := make([][]int, rows)

	// Initialize those 10 empty slices
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
	// fmt.Println(distance)

	// Iterate over the matrix to compute the cost of deletions,insertions and/or substitutions
	for col := 1; col < cols; col++ {
		for row := 1; row < rows; row++ {
			if s[row-1] == t[col-1] {
				cost = 0 // If the characters are the same in the two strings in a given position [i,j] then the cost is 0
			} else {
				// In order to align the results with those of the Python Levenshtein package, if we choose to calculate the ratio
				// the cost of a substitution is 2. If we calculate just distance, then the cost of a substitution is 1.
				if ratio_calc == true {
					cost = 2
				} else {
					cost = 1
				}
				distance[row][col] = MinOf(distance[row-1][col]+1, // Cost of deletions
					distance[row][col-1]+1,      // Cost of insertions
					distance[row-1][col-1]+cost) // Cost of substitutions
			}
		}
	}
	fmt.Println(distance)
    fmt.Println(row)
    fmt.Println(col)        

	if ratio_calc == true {
		// Computation of the Levenshtein Distance Ratio
		ratio := ((len(s) + len(t)) - distance[row][col]) / (len(s) + len(t))
		return strconv.Itoa(ratio)
	} else {
		// print(distance) # Uncomment if you want to see the matrix showing how the algorithm computes the cost of deletions,
		// insertions and/or substitutions
		// This is the minimum number of edits needed to convert string a to string b
		return fmt.Sprintf("The strings are %d edits away", distance[row][col])
	}
}

func main() {
	fmt.Println(levenshtein_ratio_and_distance("ah", "bedla", false))
}
