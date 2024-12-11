package models

import (
	"fmt"
)

// GameState represents the state of a game, including the game type, players, and the current state.
type GameState struct {
	GameType string      // Type of the game (e.g., checkers, chess)
	Players  []string    // List of player IDs or names
	State    interface{} // The actual game state (e.g., board, positions, etc.)
}

// NewGameState creates a new GameState for a given game type and players.
func NewGameState(gameType string, players []string, initialState interface{}) *GameState {
	return &GameState{
		GameType: gameType,
		Players:  players,
		State:    initialState,
	}
}

// UpdateState updates the state of the game.
func (gs *GameState) UpdateState(newState interface{}) {
	gs.State = newState
}

// GetState returns the current state of the game.
func (gs *GameState) GetState() interface{} {
	return gs.State
}

// String returns a string representation of the game state.
func (gs *GameState) String() string {
	return fmt.Sprintf("Game Type: %s, Players: %v, State: %v", gs.GameType, gs.Players, gs.State)
}
