package api

import (
	"encoding/json"
	"net/http"
)

func GameHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Message string `json:"message"`
	}

	resp := Response{Message: "Game handler working!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
