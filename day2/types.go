package day2

type (
	Game struct {
		ID   int
		Sets []Set
	}

	Set struct {
		Red   int
		Blue  int
		Green int
	}
)
