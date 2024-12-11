package tictactoe

import "testing"

func TestNewTicTacToeGame(t *testing.T) {
	game := NewGame()
	if len(game.Board) != 3 {
		t.Errorf("Expected board size of 3, got %d", len(game.Board))
	}

	// Check initial turn is "X"
	if game.Turn != "X" {
		t.Errorf("Expected initial turn to be 'X', got %s", game.Turn)
	}
}

func TestMove(t *testing.T) {
	game := NewGame()

	// Valid move (X to A1)
	err := game.Move(0, 0) // A1 -> (0, 0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid move (X tries to play on an occupied cell)
	err = game.Move(0, 0)
	if err == nil {
		t.Errorf("Expected error for invalid move, got nil")
	}
}

func TestTurnChange(t *testing.T) {
	game := NewGame()

	// First move by X
	game.Move(0, 0)
	if game.Turn != "O" {
		t.Errorf("Expected turn to be 'O', got %s", game.Turn)
	}
}

func TestWinner(t *testing.T) {
	game := NewGame()

	// Make moves to create a winner
	game.Move(0, 0) // X -> A1
	game.Move(0, 1) // O -> A2
	game.Move(1, 0) // X -> B1
	game.Move(1, 1) // O -> B2
	game.Move(2, 0) // X -> C1 (X wins)

	if winner := game.CheckWinner(); winner != "X" {
		t.Errorf("Expected winner 'X', got %s", winner)
	}
}

func TestDraw(t *testing.T) {
	game := NewGame()

	// Create a draw scenario
	game.Move(0, 0) // X -> A1
	game.Move(0, 1) // O -> A2
	game.Move(0, 2) // X -> A3
	game.Move(1, 1) // O -> B2
	game.Move(2, 0) // X -> C1
	game.Move(2, 2) // O -> C3
	game.Move(1, 0) // X -> B1
	game.Move(1, 2) // O -> B3
	game.Move(2, 1) // X -> C2 (draw)

	if winner := game.CheckWinner(); winner != "" {
		t.Errorf("Expected no winner, got %s", winner)
	}
}
