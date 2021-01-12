package match

import (
	"log"
	"regexp"
	"strings"
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

func removeUnusedChars(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, " ")
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
