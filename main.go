package main

import (
	"aoc-1/day1"
	_ "embed"
	"strings"
)

//go:embed day1/input.txt
var input string

func main() {
	// day 1 - answer is 53268
	splitInput := strings.Split(input, "\n")
	day1.CalculateSum(splitInput)
}
