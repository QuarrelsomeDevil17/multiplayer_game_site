package chess

type ChessGame struct {
	Board [][]string
}

func NewChessGame() *ChessGame {
	return &ChessGame{
		Board: initializeBoard(),
	}
}

func initializeBoard() [][]string {
	// Initialize a default chess board
	return [][]string{
		{"r", "n", "b", "q", "k", "b", "n", "r"},
		{"p", "p", "p", "p", "p", "p", "p", "p"},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"P", "P", "P", "P", "P", "P", "P", "P"},
		{"R", "N", "B", "Q", "K", "B", "N", "R"},
	}
}
