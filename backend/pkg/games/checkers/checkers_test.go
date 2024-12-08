package checkers

import "testing"

func TestNewCheckersGame(t *testing.T) {
	game := NewCheckersGame()
	if len(game.Board) != 8 {
		t.Errorf("Expected board size of 8, got %d", len(game.Board))
	}
}
