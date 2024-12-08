# Multiplayer Games Backend

This is the backend for a multiplayer games platform, supporting games like Chess, Checkers, and Tic-Tac-Toe. It uses Go for performance and scalability.

## Features
- Game management with session-based containers.
- REST and WebSocket API for communication.
- Support for Docker and Kubernetes orchestration.

## Structure
- **cmd/**: Entry point for the application.
- **pkg/**: Contains the core game logic, handlers, and utilities.
- **configs/**: Configuration files for Docker and Kubernetes.
- **scripts/**: Deployment and cleanup scripts.

## Getting Started
1. Install Go 1.21 or later.
2. Run `go mod tidy` to install dependencies.
3. Start the server:
   ```bash
   go run cmd/main.go
