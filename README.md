# go_fuzzymatch

## description:

Fuzzy matching CLI tool written in Go. The matching calculation is done using [Levenshtein distance algorithm](https://en.wikipedia.org/wiki/Levenshtein_distance). 

This tool implicitely preformats the input strings using these actions:

- applies lowercase

- removes non-alphabetical chars and non-digits

- in case the input string can be split to an array of words by whitespace, the array items are unique  

This tool is able to operate in two modes:

- `simple` mode

- `deep dive` mode

>Simple mode is faster but less precise when comparing multi-word strings.

Example results in `simple` mode:
|string 1  |string 2  |calculated output  |
|---|---|---|
|apple inc|apple inc|1|
|apple inc|Apple Inc.|1|
|Apple|Apple Inc.|1|
|Apple Inc|Apple|1|
|aplle|Apple|0.8|
|Apple Corp.|Apple Corp. GMBH|0.8|
|GMBH Apple Corp|Apple Inc.|0.5833333|
|apple Inc.|GMBH Apple Corp.|0.5833333|
|aplle Inc.|GMBH Apple Corp.|0.5|


> Deep dive mode is slower but goes more in depth when comparing multi-word strings using permutations.

Example results in `deepDive` mode:
|string 1  |string 2  |calculated output  |
|---|---|---|
|apple inc|apple inc|1|
|apple inc|Apple Inc.|1|
|Apple|Apple Inc.|1|
|Apple Inc|Apple|1|
|aplle|Apple|0.8|
|Apple Corp.|Apple Corp. GMBH|**1**|
|GMBH Apple Corp|Apple Inc.|**0.6315789**|
|apple Inc.|GMBH Apple Corp.|**0.6315789**|
|aplle Inc.|GMBH Apple Corp.|**0.5263158**|

## how to run:

`go run fuzzy_match.go -string1="apple" -string2="Apple inc." -deepDive=true|false`