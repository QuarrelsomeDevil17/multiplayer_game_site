package models

type GameState struct {
	GameType string
	Players  []string
	State    interface{}
}
