package checkers

import "testing"

func TestNewCheckersGame(t *testing.T) {
	game := NewGame()
	if len(game.Board) != 8 {
		t.Errorf("Expected board size of 8, got %d", len(game.Board))
	}

	// Check initial turn is red
	if game.Turn != "red" {
		t.Errorf("Expected initial turn to be 'red', got %s", game.Turn)
	}
}

func TestMove(t *testing.T) {
	game := NewGame()

	// Valid move (red piece from B3 to C4)
	err := game.Move("B3", "C4")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid move (no piece at D5)
	err = game.Move("D5", "E6")
	if err == nil {
		t.Errorf("Expected error for invalid move, got nil")
	}
}

func TestTurnChange(t *testing.T) {
	game := NewGame()

	// First move by red
	game.Move("B3", "C4")
	if game.Turn != "black" {
		t.Errorf("Expected turn to be 'black', got %s", game.Turn)
	}
}

func TestInvalidMove(t *testing.T) {
	game := NewGame()

	// Try moving a piece that doesn't exist
	err := game.Move("A1", "B2")
	if err == nil {
		t.Errorf("Expected error for invalid move, got nil")
	}
}
