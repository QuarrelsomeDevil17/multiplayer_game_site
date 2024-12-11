// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"backend/pkg/api"
// 	"backend/pkg/utils"
// )

// func main() {
// 	utils.LoadConfig()

// 	r := mux.NewRouter()

// 	gameHandler := api.NewGameHandler()
// 	r.HandleFunc("/api/games", gameHandler.ListGames).Methods("GET")
// 	r.HandleFunc("/api/session", gameHandler.CreateSession).Methods("POST")

// 	log.Printf("Starting server on port %d...", utils.AppConfig.Server.Port)
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }



package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"backend/pkg/api"
	"backend/pkg/container"
	"backend/pkg/models"
	"backend/pkg/utils"
)

func main() {
	// Determine config path from environment variable or use default
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "/home/qdev/github/multiplayer_game_site/backend/configs/config.yaml" // Default config file path
	}

	// Load the app's configuration
	utils.LoadConfig(configPath)

	// Initialize the Docker manager for container management if enabled
	var dockerManager *container.DockerManager
	if utils.AppConfig.Docker.Enable {
		var err error
		dockerManager, err = container.NewDockerManager()
		if err != nil {
			log.Fatalf("Error initializing Docker manager: %v", err)
		}
		log.Println("Docker manager initialized")
	} else {
		log.Println("Docker manager is disabled")
	}

	// Initialize a new router
	r := mux.NewRouter()

	// Initialize the game handler (handles game-related HTTP requests)
	gameHandler := api.NewGameHandler(dockerManager)
	r.HandleFunc("/api/games", gameHandler.ListGames).Methods("GET")
	r.HandleFunc("/api/session", gameHandler.CreateSession).Methods("POST")

	// WebSocket handler (example, placeholder for game-specific logic)
	r.HandleFunc("/ws", api.WebSocketHandler).Methods("GET")

	// Example: Initialize some in-memory sessions or games
	gameState := models.NewGameState("checkers", []string{"Player1", "Player2"}, nil) // Placeholder state
	session := models.NewSession("session_123", "checkers", "Player1", "Player2")
	log.Println("Initialized session:", session)
	log.Println("Initialized game state:", gameState)

	// Start the server
	port := utils.AppConfig.Server.Port
	if port == 0 {
		port = 8080 // Default port fallback
	}
	address := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on port %d...", port)
	log.Fatal(http.ListenAndServe(address, r))
}

