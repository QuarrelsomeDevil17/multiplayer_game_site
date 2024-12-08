package tictactoe

import "testing"

func TestNewTicTacToeGame(t *testing.T) {
	game := NewTicTacToeGame()
	if game.Board[0][0] != "" {
		t.Errorf("Expected an empty board")
	}
}
