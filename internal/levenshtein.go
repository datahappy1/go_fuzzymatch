package match

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

// LevenshteinRatio returns float32
func LevenshteinRatio(s1 string, s2 string) float32 {
	var rowLength = len(s1)
	var colLength = len(s2)
	var rows = rowLength + 1
	var cols = colLength + 1
	var cost = 0

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
