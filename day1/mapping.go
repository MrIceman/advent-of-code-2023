package day1

import (
	"fmt"
	"slices"
	"strconv"
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

func extractNumbers(s string) int {
	leftValue := leftSearch(s)
	rightValue := rightSearch(s)
	i, _ := strconv.Atoi(fmt.Sprintf("%d%d", leftValue, rightValue))

	return i
}

func leftSearch(s string) int {
	return search(s, 0, func(i int) bool {
		return i == len(s)-1
	}, func(i int) int {
		return i + 1
	})
}

func rightSearch(s string) int {
	return search(s, len(s)-1, func(i int) bool {
		return i == 0
	}, func(i int) int {
		return i - 1
	})
}

func search(s string, start int, stopFn func(int) bool, nextFn func(int) int) int {
	found := -1
	var strBuffer []string
	splitString := strings.Split(s, "")
	stopped := false
	strIdx := start
	for !stopped {
		stopped = stopFn(strIdx)
		if n, ok := isNumber(splitString[strIdx]); ok {
			found = n
			return found
		}
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
						found = slices.Index(numberList, resultString) + 1
						return found
					}
					hitIndex += 1
				} else {
					// chars don't match, this number can not be included
					strBuffer = []string{}
					continue numbersLoop
				}
			}
		}
		strIdx = nextFn(strIdx)
	}

	return found
}

func isNumber(s string) (int, bool) {
	n, err := strconv.Atoi(s)
	return n, err == nil
}
