package main

import (
	"flag"
	"fmt"
	"github.com/datahappy1/go_fuzzymatch/internal/go_fuzzymatch"
)

func main() {
	string1Ptr := flag.String("string1", "", "first string")
	string2Ptr := flag.String("string2", "", "second string")
	mode := flag.String("mode", "simple", "mode < simple | deepDive | combined >")
	flag.Parse()

	if *string1Ptr == *string2Ptr {
		fmt.Println(100)
	} else if *string1Ptr == "" || *string2Ptr == "" {
		fmt.Println(0)
	} else {
		var m = &go_fuzzymatch.Match{}
		if *mode == "simple" {
			m.Strategy = go_fuzzymatch.Simple{}
			fmt.Println(m.MatchStrings(*string1Ptr, *string2Ptr))
		} else if *mode == "deepDive" {
			m.Strategy = go_fuzzymatch.DeepDive{}
			fmt.Println(m.MatchStrings(*string1Ptr, *string2Ptr))
		} else if *mode == "combined" {
			m.Strategy = go_fuzzymatch.Combined{}
			fmt.Println(m.MatchStrings(*string1Ptr, *string2Ptr))
		} else {
			fmt.Println("Unknown mode argument value, options are: simple, deepDive, combined")
		}
	}
}
