package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	splitInput := strings.Split(input, "\n")
	log.Printf("elements: %d", len(splitInput))
	sum := 0
	for _, s := range splitInput {
		sum += parseToInt(s)
	}
	log.Printf("total sum of all the calibration values: %d", sum)
}

func parseToInt(s string) int {
	// a map which holds the indexes of each found numbers, the result is then a concatenation of the first and the last found number
	indexMap := map[int]int{}

	for idx, c := range strings.Split(s, "") {
		if n, ok := isNumber(c); ok {
			indexMap[idx] = n
		}
	}

	if len(indexMap) == 0 {
		return 0
	}

	if len(indexMap) == 1 {
		return parseStringNumberPairToInt(indexMap[0], indexMap[0])
	}
	return parseStringNumberPairToInt(getLowestAndHighestValuesFromIndexMap(indexMap))
}

func isNumber(s string) (int, bool) {
	n, err := strconv.Atoi(s)
	return n, err == nil

}

func parseStringNumberPairToInt(x, y int) int {
	res, _ := strconv.Atoi(fmt.Sprintf("%d%d", x, y))

	return res
}

func getLowestAndHighestValuesFromIndexMap(indexMap map[int]int) (lowestValue, highestValue int) {
	// initialize both values first to the first index
	lowestKey, highestKey := 0, 0
	for k, v := range indexMap {
		if v < lowestValue {
			lowestKey = k
			lowestValue = v
		}
		if v > highestValue {
			highestKey = k
			highestValue = k
		}
	}
	return lowestKey, highestKey
}
