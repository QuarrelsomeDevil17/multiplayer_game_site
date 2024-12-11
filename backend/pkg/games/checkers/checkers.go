package checkers

import (
	"fmt"
	"math"
)

// CheckersBoard structure to hold the state of the board
type CheckersBoard struct {
	Board [8][8]string // 8x8 checkers board
	Turn  string       // "red" or "black"
}

// NewGame initializes a new Checkers game
func NewGame() *CheckersBoard {
	board := [8][8]string{
		{"r", " ", "r", " ", "r", " ", "r", " "},
		{" ", "r", " ", "r", " ", "r", " ", "r"},
		{"r", " ", "r", " ", "r", " ", "r", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{"b", " ", "b", " ", "b", " ", "b", " "},
		{" ", "b", " ", "b", " ", "b", " ", "b"},
		{"b", " ", "b", " ", "b", " ", "b", " "},
	}
	return &CheckersBoard{Board: board, Turn: "red"}
}

// Move executes a move on the board
func (c *CheckersBoard) Move(from, to string) error {
	fromRow, fromCol := parsePosition(from)
	toRow, toCol := parsePosition(to)

	// Validate the move
	if c.Board[fromRow][fromCol] == "" {
		return fmt.Errorf("no piece at starting position")
	}

	if !isValidMove(c, fromRow, fromCol, toRow, toCol) {
		return fmt.Errorf("invalid move")
	}

	// Perform the move
	c.Board[toRow][toCol] = c.Board[fromRow][fromCol]
	c.Board[fromRow][fromCol] = ""

	// Promote to king if piece reaches the opponent's end row
	if (toRow == 0 && c.Turn == "red") || (toRow == 7 && c.Turn == "black") {
		c.Board[toRow][toCol] = string(c.Board[toRow][toCol][0]) + "K" // Add "K" to mark as king
	}

	// Change turn
	c.Turn = toggleTurn(c.Turn)
	return nil
}

// parsePosition converts a position like "A1" to row and column
func parsePosition(pos string) (int, int) {
	row := 8 - int(pos[1]-'0')
	col := int(pos[0] - 'A')
	return row, col
}

// isValidMove checks if a move is valid
func isValidMove(c *CheckersBoard, fromRow, fromCol, toRow, toCol int) bool {
	piece := c.Board[fromRow][fromCol]
	if (c.Turn == "red" && piece[0] != 'r') || (c.Turn == "black" && piece[0] != 'b') {
		return false // Moving opponent's piece
	}

	rowDiff := toRow - fromRow
	colDiff := toCol - fromCol

	// Regular pieces move forward diagonally
	if !isKing(piece) {
		if (c.Turn == "red" && rowDiff >= 0) || (c.Turn == "black" && rowDiff <= 0) {
			return false // Must move forward
		}

		if math.Abs(float64(colDiff)) != 1 || math.Abs(float64(rowDiff)) != 1 {
			// Invalid simple diagonal move
			return false
		}

		if c.Board[toRow][toCol] != "" {
			return false // Destination must be empty
		}
	} else { // King pieces can move in any diagonal direction
		if math.Abs(float64(rowDiff)) != 1 || math.Abs(float64(colDiff)) != 1 {
			return false // Invalid diagonal move
		}

		if c.Board[toRow][toCol] != "" {
			return false // Destination must be empty
		}
	}

	// Implement capture logic
	if math.Abs(float64(rowDiff)) == 2 && math.Abs(float64(colDiff)) == 2 {
		midRow := (fromRow + toRow) / 2
		midCol := (fromCol + toCol) / 2
		capturedPiece := c.Board[midRow][midCol]
		if capturedPiece == "" || capturedPiece[0] == piece[0] {
			return false // No opponent's piece to capture
		}
		// Remove the captured piece
		c.Board[midRow][midCol] = ""
	}

	return true
}

// isKing checks if the piece is a king
func isKing(piece string) bool {
	return len(piece) > 1 && piece[1] == 'K'
}

// toggleTurn switches turns between "red" and "black"
func toggleTurn(turn string) string {
	if turn == "red" {
		return "black"
	}
	return "red"
}

