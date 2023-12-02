package day2

import (
	"log"
	"testing"
)

func TestShouldMapInputToExpectedStruct(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expected := []Game{
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

	result := MustMapInputToGames(input)

	for i, g := range result {
		e := expected[i]
		if e.ID != g.ID {
			log.Printf("expected game ID %d but got %d", e, g.ID)
			t.Fail()
		}
		for j, s := range g.Sets {
			if s.Blue != e.Sets[j].Blue {
				log.Printf("%d: expected set %d blue %d but got %d", g.ID, j, e.Sets[j].Blue, s.Blue)
				t.Fail()
			}
			if s.Red != e.Sets[j].Red {
				log.Printf("%d: expected set %d red %d but got %d", g.ID, j, e.Sets[j].Red, s.Red)
				t.Fail()
			}
			if s.Green != e.Sets[j].Green {
				log.Printf("%d: expected set %d green %d but got %d", g.ID, j, e.Sets[j].Green, s.Green)
				t.Fail()
			}
		}
	}
}
