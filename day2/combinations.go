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
