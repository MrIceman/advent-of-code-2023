package day2

func GetResultPart1(input []string) int {
	games := MustMapInputToGames(input)
	_, result := GetPossibleGames(games, 12, 13, 14)

	return result
}
