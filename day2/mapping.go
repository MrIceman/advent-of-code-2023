package day2

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func MustMapInputToGames(s []string) []Game {
	var games []Game
	for _, v := range s {
		g, err := parseGame(v)
		if err != nil {
			log.Printf("Oops, a game could not be parsed!")
			panic(err)
		}
		games = append(games, g)
	}
	return games
}

func parseGame(input string) (Game, error) {
	var game Game
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 {
		return game, fmt.Errorf("invalid format")
	}

	id, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(parts[0], "Game")))
	if err != nil {
		return game, fmt.Errorf("invalid game ID")
	}
	game.ID = id

	setStrings := strings.Split(parts[1], ";")
	for _, setStr := range setStrings {
		set, err := parseSet(strings.TrimSpace(setStr))
		if err != nil {
			return game, err
		}
		game.Sets = append(game.Sets, set)
	}

	return game, nil
}

func parseSet(input string) (Set, error) {
	var set Set
	colorCounts := regexp.MustCompile(`(\d+)\s*(\w+)`).FindAllStringSubmatch(input, -1)
	for _, colorCount := range colorCounts {
		count, _ := strconv.Atoi(colorCount[1])
		switch strings.ToLower(colorCount[2]) {
		case "red":
			set.Red += count
		case "blue":
			set.Blue += count
		case "green":
			set.Green += count
		default:
			return set, fmt.Errorf("unknown color: %s", colorCount[2])
		}
	}
	return set, nil
}
