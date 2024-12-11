package models

import "fmt"

// Session represents a game session between two players.
type Session struct {
	ID       string // Unique ID for the session
	GameType string // Type of the game (e.g., checkers, chess)
	Player1  string // Player 1's name or ID
	Player2  string // Player 2's name or ID
}

// NewSession creates a new session with a unique session ID, game type, and two players.
func NewSession(id, gameType, player1, player2 string) *Session {
	return &Session{
		ID:       id,
		GameType: gameType,
		Player1:  player1,
		Player2:  player2,
	}
}

// String returns a string representation of the session.
func (s *Session) String() string {
	return fmt.Sprintf("Session ID: %s, Game Type: %s, Players: %s vs %s", s.ID, s.GameType, s.Player1, s.Player2)
}
