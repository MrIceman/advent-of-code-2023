package main

import (
	"slices"
	"strings"
)

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

var numberList = []string{
	one,
	two,
	three,
	four,
	five,
	six,
	seven,
	eight,
	nine,
}

func extractNumbersFromString(s string) map[int]int {
	result := map[int]int{}
	var strBuffer []string
	splitString := strings.Split(s, "")
strLoop:
	for strIdx := range splitString {
	numbersLoop:
		for _, number := range numberList {
			hitIndex := strIdx
			// we reached the end
			if hitIndex > len(splitString)-3 {
				break numbersLoop
			}
			for _, n := range strings.Split(number, "") {
				if splitString[hitIndex] == n {
					strBuffer = append(strBuffer, splitString[hitIndex])
					resultString := strings.Join(strBuffer, "")
					if resultString == number {
						// the numbers are ordered so we can safely assume that the
						// index of the number in the list is equivalent to the value - 1
						result[strIdx] = slices.Index(numberList, resultString) + 1
						continue strLoop
					}
					hitIndex += 1
				} else {
					// chars don't match, this number can not be included
					strBuffer = []string{}
					continue numbersLoop
				}
			}
		}
	}

	return result
}
