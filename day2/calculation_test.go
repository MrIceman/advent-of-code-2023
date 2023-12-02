package day2

import (
	"log"
	"testing"
)

func TestShouldReturnExpectedPossibleGames(t *testing.T) {
	input := []Game{
		{
			ID: 1,
			Sets: []Set{
				{
					Blue: 3,
					Red:  4,
				},
				{
					Red:   1,
					Green: 2,
					Blue:  6,
				},
				{
					Green: 2,
				},
			},
		},
		{
			ID: 2,
			Sets: []Set{
				{
					Blue:  1,
					Green: 2,
				},
				{
					Green: 3,
					Blue:  4,
					Red:   1,
				},
				{
					Green: 1,
					Blue:  1,
				},
			},
		},
		{
			ID: 3,
			Sets: []Set{
				{
					Green: 8,
					Blue:  6,
					Red:   20,
				},
				{
					Blue:  5,
					Red:   4,
					Green: 13,
				},
				{
					Green: 5,
					Red:   1,
				},
			},
		},
		{
			ID: 4,
			Sets: []Set{
				{
					Green: 1,
					Red:   3,
					Blue:  6,
				},
				{
					Green: 3,
					Red:   6,
				},
				{
					Green: 3,
					Blue:  15,
					Red:   14,
				},
			},
		},
		{
			ID: 5,
			Sets: []Set{
				{
					Red:   6,
					Blue:  1,
					Green: 3,
				},
				{
					Blue:  2,
					Red:   1,
					Green: 2,
				},
			},
		},
	}

	result, sum := GetPossibleGames(input, 12, 13, 14)

	if len(result) != 3 {
		log.Printf("expected to have 3 games but got %d", len(result))
		t.Fail()
	}
	if sum != 8 {
		t.Fail()
	}
}

func TestShouldGetExpectedAccumulatedSum(t *testing.T) {
	input := []Game{
		{
			ID: 1,
			Sets: []Set{
				{
					Blue: 3,
					Red:  4,
				},
				{
					Red:   1,
					Green: 2,
					Blue:  6,
				},
				{
					Green: 2,
				},
			},
		},
		{
			ID: 2,
			Sets: []Set{
				{
					Blue:  1,
					Green: 2,
				},
				{
					Green: 3,
					Blue:  4,
					Red:   1,
				},
				{
					Green: 1,
					Blue:  1,
				},
			},
		},
		{
			ID: 3,
			Sets: []Set{
				{
					Green: 8,
					Blue:  6,
					Red:   20,
				},
				{
					Blue:  5,
					Red:   4,
					Green: 13,
				},
				{
					Green: 5,
					Red:   1,
				},
			},
		},
		{
			ID: 4,
			Sets: []Set{
				{
					Green: 1,
					Red:   3,
					Blue:  6,
				},
				{
					Green: 3,
					Red:   6,
				},
				{
					Green: 3,
					Blue:  15,
					Red:   14,
				},
			},
		},
		{
			ID: 5,
			Sets: []Set{
				{
					Red:   6,
					Blue:  1,
					Green: 3,
				},
				{
					Blue:  2,
					Red:   1,
					Green: 2,
				},
			},
		},
	}

	result := GetAddedPowerOfSets(input)

	if result != 2286 {
		log.Printf("expected result to be %d but got %d", 2286, result)
		t.Fail()
	}
}
