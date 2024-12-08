package chess

import "testing"

func TestNewChessGame(t *testing.T) {
	game := NewChessGame()
	if len(game.Board) != 8 {
		t.Errorf("Expected board size of 8, got %d", len(game.Board))
	}
}
