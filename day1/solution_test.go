package day1

import (
	"log"
	"testing"
)

func TestCalculateSum(t *testing.T) {
	i := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := 281

	sum := CalculateSum(i)

	if sum != expected {
		log.Printf("expected result to be %d but it was %d", expected, sum)
		t.Fail()
	}
}
