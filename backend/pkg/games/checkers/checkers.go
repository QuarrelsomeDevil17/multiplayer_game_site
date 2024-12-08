package checkers

type CheckersGame struct {
	Board [][]string
}

func NewCheckersGame() *CheckersGame {
	return &CheckersGame{
		Board: initializeBoard(),
	}
}

func initializeBoard() [][]string {
	// Default Checkers board setup
	return [][]string{
		{"b", "", "b", "", "b", "", "b", ""},
		{"", "b", "", "b", "", "b", "", "b"},
		{"b", "", "b", "", "b", "", "b", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "r", "", "r", "", "r", "", "r"},
		{"r", "", "r", "", "r", "", "r", ""},
		{"", "r", "", "r", "", "r", "", "r"},
	}
}
