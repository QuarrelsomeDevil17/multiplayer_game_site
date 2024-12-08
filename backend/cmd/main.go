package main

import (
	"log"
	"net/http"

	"backend/pkg/api"
	"backend/pkg/utils"
)

func main() {
	utils.InitializeLogger()
	http.HandleFunc("/api/games", api.GameHandler)
	http.HandleFunc("/api/ws", api.WebSocketHandler)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
