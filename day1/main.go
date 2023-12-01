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
	var ints []string

	for _, c := range strings.Split(s, "") {
		if ok := isNumber(c); ok {
			ints = append(ints, c)
		}
	}

	if len(ints) == 0 {
		return 0
	}

	if len(ints) == 1 {
		return parseStringNumberPairToInt(ints[0], ints[0])
	}
	return parseStringNumberPairToInt(ints[0], ints[len(ints)-1])
}

func parseStringNumberPairToInt(x, y string) int {
	res, _ := strconv.Atoi(fmt.Sprintf("%s%s", x, y))

	return res
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil

}
