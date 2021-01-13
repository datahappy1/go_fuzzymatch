# go_fuzzymatch

## description:

Fuzzy matching CLI tool written in Go. The matching calculation is done using [Levenshtein distance algorithm](https://en.wikipedia.org/wiki/Levenshtein_distance). 

This tool implicitly formats the input strings using these actions:

- applies lowercase
- processes only alphabetical chars and digits
- in case the input string can be split by whitespace into an array of words, the array items are deduplicated  

This tool is able to operate in two modes:
- `simple` mode
- `deepDive` mode
- `combined` mode

### Simple mode:
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


### DeepDive mode:
> DeepDive mode is slower but goes more in depth when comparing multi-word strings using permutations. If DeepDive mode
>is provided with two single word strings for comparison, it behaves exactly like the simple mode.

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

### Combined mode:
> Combined mode starts with simple mode and if no match above `0.85`, goes for a deepDive mode.


## how to run:

`go run fuzzy_match.go -string1="apple" -string2="Apple inc." -mode=simple|deepDive|combined`