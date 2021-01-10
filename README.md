# go_fuzzymatch

## description:
Fuzzy matching CLI tool written in Go. The matching calculation is based on a calculated ratio of [Levenshtein distance algorithm](https://en.wikipedia.org/wiki/Levenshtein_distance). 

This tool implicitely preformats the input strings using these actions:
- applies lowercase
- removes non-alphabetical chars and non-digits

This tool is able to operate in two modes:
- `simple` mode
- `deep dive` mode

Simple mode is faster but less precise when comparing multi-word strings. Deep dive mode is slower but goes more in depth when comparing multi-word strings using permutations.

## how to run:
`go run fuzzy_match.go -string1="apple" -string2="Apple inc." -deepDive=true|false`