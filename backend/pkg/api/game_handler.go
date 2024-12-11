package api

import (
	"encoding/json"
	"net/http"
	"backend/pkg/container"
)

// GameHandler handles game-related HTTP requests
type GameHandler struct {
    DockerManager *container.DockerManager
}

// NewGameHandler initializes a new GameHandler
func NewGameHandler(dockerManager *container.DockerManager) *GameHandler {
    return &GameHandler{
        DockerManager: dockerManager,
    }
}

func (h *GameHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Session created"}
	json.NewEncoder(w).Encode(response)
}

func (h *GameHandler) ListGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"games": []string{"chess", "checkers", "tictactoe"},
	}
	json.NewEncoder(w).Encode(response)
}
