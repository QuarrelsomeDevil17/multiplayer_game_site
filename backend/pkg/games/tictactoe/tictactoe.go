package tictactoe

type TicTacToeGame struct {
	Board [3][3]string
}

func NewTicTacToeGame() *TicTacToeGame {
	return &TicTacToeGame{}
}

func (g *TicTacToeGame) MakeMove(player string, row, col int) bool {
	if g.Board[row][col] == "" {
		g.Board[row][col] = player
		return true
	}
	return false
}
