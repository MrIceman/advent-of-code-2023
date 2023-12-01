package main

import (
	"aoc-1/day1"
	_ "embed"
	"strings"
)

//go:embed day1/input.txt
var input string

func main() {
	splitInput := strings.Split(input, "\n")
	day1.CalculateSum(splitInput)
}
