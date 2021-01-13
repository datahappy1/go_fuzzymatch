package main

import (
	"flag"
	"fmt"

	match "./internal"
)

func main() {
	string1Ptr := flag.String("string1", "", "first string")
	string2Ptr := flag.String("string2", "", "second string")
	deepDive := flag.Bool("deepDive", false, "deep dive")
	flag.Parse()

	if *string1Ptr == *string2Ptr {
		fmt.Println(1)
	} else if *string1Ptr == "" || *string2Ptr == "" {
		fmt.Println(0)
	} else {
		var m = &match.Match{}
		if *deepDive == true {
			m.Strategy = match.DeepDive{}
			fmt.Println(m.MatchStrings(*string1Ptr, *string2Ptr))
		} else {
			m.Strategy = match.Simple{}
			fmt.Println(m.MatchStrings(*string1Ptr, *string2Ptr))
		}
	}
}
