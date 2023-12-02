package day2

func GetPossibleGames(input []Game, red, green, blue int) ([]Game, int) {
	var possibleGames []Game
	idSum := 0
	for _, g := range input {
		isPossible := true
	setLoop:
		for _, s := range g.Sets {
			if red < s.Red || blue < s.Blue || green < s.Green {
				isPossible = false
				break setLoop
			}
		}
		if isPossible {
			idSum += g.ID
		}
	}
	return possibleGames, idSum
}

// GetAddedPowerOfSets is part 2 of the task
func GetAddedPowerOfSets(input []Game) int {
	powerSum := 0
	for _, g := range input {
		rMax, gMax, bMax := 0, 0, 0
		for _, s := range g.Sets {
			if s.Red > rMax {
				rMax = s.Red
			}
			if s.Blue > bMax {
				bMax = s.Blue
			}
			if s.Green > gMax {
				gMax = s.Green
			}
		}
		powerSum += rMax * gMax * bMax
	}
	return powerSum
}
