package tictactoe

import (
	"fmt"
)

// TicTacToeBoard structure to hold the state of the board
type TicTacToeBoard struct {
	Board [3][3]string // 3x3 board
	Turn  string       // "X" or "O"
}

// NewGame initializes a new Tic-Tac-Toe game
func NewGame() *TicTacToeBoard {
	board := [3][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	return &TicTacToeBoard{Board: board, Turn: "X"}
}

// Move executes a move on the board
func (t *TicTacToeBoard) Move(row, col int) error {
	// Validate the move
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return fmt.Errorf("move out of bounds")
	}

	if t.Board[row][col] != "" {
		return fmt.Errorf("cell already occupied")
	}

	// Perform the move
	t.Board[row][col] = t.Turn

	// Check for a winner
	if winner := t.CheckWinner(); winner != "" {
		return fmt.Errorf("player %s wins", winner)
	}

	// Change turn
	if t.Turn == "X" {
		t.Turn = "O"
	} else {
		t.Turn = "X"
	}
	return nil
}

// CheckWinner checks the board for a winner
func (t *TicTacToeBoard) CheckWinner() string {
	// Check rows and columns for a winner
	for i := 0; i < 3; i++ {
		if t.Board[i][0] == t.Board[i][1] && t.Board[i][1] == t.Board[i][2] && t.Board[i][0] != "" {
			return t.Board[i][0]
		}
		if t.Board[0][i] == t.Board[1][i] && t.Board[1][i] == t.Board[2][i] && t.Board[0][i] != "" {
			return t.Board[0][i]
		}
	}

	// Check diagonals for a winner
	if t.Board[0][0] == t.Board[1][1] && t.Board[1][1] == t.Board[2][2] && t.Board[0][0] != "" {
		return t.Board[0][0]
	}
	if t.Board[0][2] == t.Board[1][1] && t.Board[1][1] == t.Board[2][0] && t.Board[0][2] != "" {
		return t.Board[0][2]
	}

	return ""
}
