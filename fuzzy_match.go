package main

import (
	"flag"
	"fmt"

	match "github.com/datahappy1/go_fuzzymatch/match_core"
)

func main() {
	string1Ptr := flag.String("string1", "apple", "first string")
	string2Ptr := flag.String("string2", "bear", "second string")
	deepDive := flag.Bool("deepDive", false, "deep dive")
	flag.Parse()

	var m = &match.Match{}
	if *deepDive == true {
		m.Strategy = match.DeepDive{}
		fmt.Println(m.MatchStrings(*string1Ptr, *string2Ptr))
	} else {
		m.Strategy = match.Simple{}
		fmt.Println(m.MatchStrings(*string1Ptr, *string2Ptr))
	}
}
