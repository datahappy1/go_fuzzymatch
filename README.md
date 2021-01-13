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
|apple inc|apple inc|1|
|apple inc|Apple Inc.|1|
|Apple|Apple Inc.|1|
|Apple Inc|Apple|1|
|aplle|Apple|0.8|
|Apple Corp.|Apple Corp. GMBH|0.8|
|GMBH Apple Corp|Apple Inc.|0.5833333|
|apple Inc.|GMBH Apple Corp.|0.5833333|
|aplle Inc.|GMBH Apple Corp.|0.5|


### deepDive mode:
> deepDive mode is slower but goes more in depth when comparing multi-word strings using permutations. 
> If deepDive mode is provided with two single word strings for comparison, it evaluates the strings like in simple mode.

Example results in `deepDive` mode:
|string 1  |string 2  |calculated output  |
|---|---|---|
|apple inc|apple inc|1|
|apple inc|Apple Inc.|1|
|Apple|Apple Inc.|1|
|Apple Inc|Apple|1|
|aplle|Apple|0.8|
|Apple Corp.|Apple Corp. GMBH|**1**|
|GMBH Apple Corp|Apple Inc.|**0.7368421**|
|apple Inc.|GMBH Apple Corp.|**0.7368421**|
|aplle Inc.|GMBH Apple Corp.|**0.6315789**|

### combined mode:
> combined mode starts strings evaluation using simple mode and if no match above `0.85` is found, it evaluates the strings again in a deepDive mode.


## how to run:

`go run fuzzy_match.go -string1="apple" -string2="Apple inc." -mode=simple|deepDive|combined`