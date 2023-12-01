package day1

import (
	"log"
	"testing"
)

func TestShouldParseStringAsExpected(t *testing.T) {
	args := []struct {
		input    string
		expected int
	}{
		{
			"oneb",
			11,
		},
		{
			"two1nine",
			29,
		},
		{
			"eighttwothree",
			83,
		},
		{
			"abcone2threexyz",
			13,
		},
		{
			"xtwone3four",
			24,
		},
		{
			"4nineeightseven2",
			42,
		},
	}

	for _, i := range args {
		log.Printf("bla")
		result := extractNumbers(i.input)
		if result != i.expected {
			log.Printf("expected %d but got %d", i.expected, result)
			t.FailNow()
		}
	}
}
