package chess

import (
	"errors"
	"fmt"
)

// ChessBoard structure to hold the state of the board
type ChessBoard struct {
	Board [8][8]string // 8x8 chessboard
	Turn  string       // "white" or "black"
}

// NewGame initializes a new Chess game
func NewGame() *ChessBoard {
	board := [8][8]string{
		{"r", "n", "b", "q", "k", "b", "n", "r"},
		{"p", "p", "p", "p", "p", "p", "p", "p"},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"P", "P", "P", "P", "P", "P", "P", "P"},
		{"R", "N", "B", "Q", "K", "B", "N", "R"},
	}
	return &ChessBoard{Board: board, Turn: "white"}
}

// Move executes a move on the board
func (c *ChessBoard) Move(from, to string) error {
	fromRow, fromCol, err := parsePosition(from)
	if err != nil {
		return err
	}

	toRow, toCol, err := parsePosition(to)
	if err != nil {
		return err
	}

	// Validate if the source position has a piece
	if c.Board[fromRow][fromCol] == "" {
		return fmt.Errorf("no piece at starting position %s", from)
	}

	// Ensure the move is valid for the specific piece
	if !isValidMove(c, fromRow, fromCol, toRow, toCol) {
		return fmt.Errorf("invalid move from %s to %s", from, to)
	}

	// Perform the move
	c.Board[toRow][toCol] = c.Board[fromRow][fromCol]
	c.Board[fromRow][fromCol] = ""

	// Change turn
	if c.Turn == "white" {
		c.Turn = "black"
	} else {
		c.Turn = "white"
	}

	return nil
}

// parsePosition converts a position like "A1" to row and column
func parsePosition(pos string) (int, int, error) {
	if len(pos) != 2 {
		return -1, -1, errors.New("invalid position format")
	}

	col := int(pos[0] - 'A')
	row := 8 - int(pos[1]-'0')

	if row < 0 || row > 7 || col < 0 || col > 7 {
		return -1, -1, errors.New("position out of bounds")
	}

	return row, col, nil
}

// isValidMove checks if a move is valid (simplified for demonstration)
func isValidMove(c *ChessBoard, fromRow, fromCol, toRow, toCol int) bool {
	piece := c.Board[fromRow][fromCol]
	destination := c.Board[toRow][toCol]

	// Check for same-team piece at the destination
	if destination != "" && (isWhite(piece) == isWhite(destination)) {
		return false
	}

	// Validate moves based on piece type
	switch piece {
	case "P": // White pawn
		return isValidPawnMove(c, fromRow, fromCol, toRow, toCol, "white")
	case "p": // Black pawn
		return isValidPawnMove(c, fromRow, fromCol, toRow, toCol, "black")
	case "R", "r": // Rook
		return isValidRookMove(c, fromRow, fromCol, toRow, toCol)
	case "N", "n": // Knight
		return isValidKnightMove(fromRow, fromCol, toRow, toCol)
	case "B", "b": // Bishop
		return isValidBishopMove(c, fromRow, fromCol, toRow, toCol)
	case "Q", "q": // Queen
		return isValidQueenMove(c, fromRow, fromCol, toRow, toCol)
	case "K", "k": // King
		return isValidKingMove(fromRow, fromCol, toRow, toCol)
	}

	return false
}

// isWhite determines if the piece is white
func isWhite(piece string) bool {
	return piece >= "A" && piece <= "Z"
}

// isValidPawnMove validates a pawn's move
func isValidPawnMove(c *ChessBoard, fromRow, fromCol, toRow, toCol int, color string) bool {
	dir := -1 // White pawns move up
	if color == "black" {
		dir = 1 // Black pawns move down
	}

	// Moving forward
	if fromCol == toCol {
		// One step forward
		if toRow == fromRow+dir && c.Board[toRow][toCol] == "" {
			return true
		}
		// Two steps forward (only from starting position)
		startRow := 6
		if color == "black" {
			startRow = 1
		}
		if fromRow == startRow && toRow == fromRow+2*dir && c.Board[toRow][toCol] == "" && c.Board[fromRow+dir][toCol] == "" {
			return true
		}
	}

	// Diagonal capture
	if abs(toCol-fromCol) == 1 && toRow == fromRow+dir && c.Board[toRow][toCol] != "" && isWhite(c.Board[toRow][toCol]) != isWhite(c.Board[fromRow][fromCol]) {
		return true
	}

	return false
}

// Utility function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// isValidRookMove validates a rook's move
func isValidRookMove(c *ChessBoard, fromRow, fromCol, toRow, toCol int) bool {
	// Rooks move in straight lines: either rows or columns must remain constant
	if fromRow != toRow && fromCol != toCol {
		return false
	}

	// Ensure no pieces block the path
	if fromRow == toRow { // Horizontal move
		start, end := min(fromCol, toCol), max(fromCol, toCol)
		for col := start + 1; col < end; col++ {
			if c.Board[fromRow][col] != "" {
				return false
			}
		}
	} else if fromCol == toCol { // Vertical move
		start, end := min(fromRow, toRow), max(fromRow, toRow)
		for row := start + 1; row < end; row++ {
			if c.Board[row][fromCol] != "" {
				return false
			}
		}
	}

	return true
}

// isValidKnightMove validates a knight's move
func isValidKnightMove(fromRow, fromCol, toRow, toCol int) bool {
	dRow, dCol := abs(toRow-fromRow), abs(toCol-fromCol)
	return (dRow == 2 && dCol == 1) || (dRow == 1 && dCol == 2)
}

// isValidBishopMove validates a bishop's move
func isValidBishopMove(c *ChessBoard, fromRow, fromCol, toRow, toCol int) bool {
	// Bishops move diagonally: the absolute difference in rows and columns must be equal
	if abs(toRow-fromRow) != abs(toCol-fromCol) {
		return false
	}

	// Ensure no pieces block the diagonal path
	rowDir := 1
	if toRow < fromRow {
		rowDir = -1
	}

	colDir := 1
	if toCol < fromCol {
		colDir = -1
	}

	for i := 1; i < abs(toRow-fromRow); i++ {
		if c.Board[fromRow+i*rowDir][fromCol+i*colDir] != "" {
			return false
		}
	}

	return true
}

// isValidQueenMove validates a queen's move
func isValidQueenMove(c *ChessBoard, fromRow, fromCol, toRow, toCol int) bool {
	// The queen can move like a rook or a bishop
	return isValidRookMove(c, fromRow, fromCol, toRow, toCol) || isValidBishopMove(c, fromRow, fromCol, toRow, toCol)
}

// Utility functions for min and max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// isValidKingMove validates a king's move
func isValidKingMove(fromRow, fromCol, toRow, toCol int) bool {
	dRow, dCol := abs(toRow-fromRow), abs(toCol-fromCol)
	return dRow <= 1 && dCol <= 1
}
