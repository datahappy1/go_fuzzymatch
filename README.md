# go_fuzzymatch
[![Go Report Card](https://goreportcard.com/badge/github.com/datahappy1/go_fuzzymatch)](https://goreportcard.com/report/github.com/datahappy1/go_fuzzymatch)
## description:

Fuzzy matching CLI tool written in Go. The string distance matching calculation is done using [Levenshtein distance algorithm](https://en.wikipedia.org/wiki/Levenshtein_distance). 

This tool implicitly formats the input strings using these actions:

- applies lowercase
- processes only alphabetical chars and digits
- in case the input string can be split by whitespace into an array of words, the array items are deduplicated  

This tool is able to operate in three modes:
- `simple` mode
- `deepDive` mode
- `combined` mode

### simple mode:
> simple mode is faster but less precise when comparing multi-word strings.

Example results in `simple` mode:
|string 1  |string 2  |calculated output  |
|---|---|---|
|apple inc|apple inc|100|
|apple inc|Apple Inc.|100|
|Apple|Apple Inc.|100|
|Apple Inc|Apple|100|
|aplle|Apple|80|
|Apple Corp.|Apple Corp. GMBH|80|
|GMBH Apple Corp|Apple Inc.|58|
|apple Inc.|GMBH Apple Corp.|58|
|aplle Inc.|GMBH Apple Corp.|50|


### deepDive mode:
> deepDive mode is slower but goes more in depth when comparing multi-word strings using permutations. 
> In this mode, if provided with two single word strings for comparison, the tool evaluates the strings like in simple mode.

Example results in `deepDive` mode:
|string 1  |string 2  |calculated output  |
|---|---|---|
|apple inc|apple inc|100|
|apple inc|Apple Inc.|100|
|Apple|Apple Inc.|100|
|Apple Inc|Apple|100|
|aplle|Apple|80|
|Apple Corp.|Apple Corp. GMBH|**100**|
|GMBH Apple Corp|Apple Inc.|**73**|
|apple Inc.|GMBH Apple Corp.|**73**|
|aplle Inc.|GMBH Apple Corp.|**63**|

### combined mode:
> combined mode starts strings evaluation using simple mode and if no match above `85` is found, it evaluates the strings again in a deepDive mode.


## how to run:
`go run cmd\fuzzy_match.go -string1="apple" -string2="Apple inc." -mode=simple|deepDive|combined`