package main

import (
	"fmt"
	"log"
	"testing"
)

func TestShouldParseStringAsExpected(t *testing.T) {
	args := []struct {
		input    string
		expected map[int]int
	}{
		{
			"oneb",
			map[int]int{
				0: 1,
			},
		},
		{
			"two1nine",
			map[int]int{
				0: 2,
				4: 9,
			},
		},
		{
			"eighttwothree",
			map[int]int{
				0: 8,
				5: 2,
				8: 3,
			},
		},
		{
			"abcone2threexyz",
			map[int]int{
				3: 1,
				7: 3,
			},
		},
		{
			"xtwone3four",
			map[int]int{
				1: 2,
				3: 1,
				7: 4,
			},
		},
		{
			"4nineeightseven2",
			map[int]int{
				1:  9,
				5:  8,
				10: 7,
			},
		},
	}

	for _, i := range args {
		log.Printf("bla")
		result := extractNumbersFromString(i.input)
		err := validateResult(result, i.expected)
		if err != nil {
			log.Printf("case \"%s\" failed: %s", i.input, err.Error())
			t.FailNow()
		}
	}
}

func validateResult(result, expected map[int]int) error {
	if len(result) != len(expected) {
		return fmt.Errorf("expected length %d but got %d", len(expected), len(result))
	}
	for k, v := range expected {
		if result[k] != v {
			return fmt.Errorf("values don't match - expected (%d, %d) but got (%d, %d)", k, v, k, result[k])
		}
	}

	return nil
}
