package chess

import "testing"

func TestNewChessGame(t *testing.T) {
	game := NewGame()
	if len(game.Board) != 8 {
		t.Errorf("Expected board size of 8, got %d", len(game.Board))
	}

	// Check initial turn is white
	if game.Turn != "white" {
		t.Errorf("Expected initial turn to be 'white', got %s", game.Turn)
	}
}

func TestMove(t *testing.T) {
	game := NewGame()

	// Valid move (white pawn from E2 to E4)
	err := game.Move("E2", "E4")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid move (no piece at A3)
	err = game.Move("A3", "A5")
	if err == nil {
		t.Errorf("Expected error for invalid move, got nil")
	}
}

func TestTurnChange(t *testing.T) {
	game := NewGame()

	// First move by white
	game.Move("E2", "E4")
	if game.Turn != "black" {
		t.Errorf("Expected turn to be 'black', got %s", game.Turn)
	}
}

func TestInvalidMove(t *testing.T) {
	game := NewGame()

	// Try moving a piece that doesn't exist
	err := game.Move("A1", "A3")
	if err == nil {
		t.Errorf("Expected error for invalid move, got nil")
	}
}
