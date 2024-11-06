package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/projeto-de-algoritmos-2024/Grafos1_GoMazing/algorithms"
)

type MazeRequest struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Algo   int `json:"algo"`
}

type SolveRequest struct {
	Algo int `json:"algo"`
}

var maze *algorithms.Maze

func generateMaze(w http.ResponseWriter, r *http.Request) {
	var req MazeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maze = algorithms.NewMaze(req.Width, req.Height)
	switch req.Algo {
	case 1:
		maze.GenerateDFS()
	case 2:
		maze.GeneratePrim()
	case 3:
		maze.GenerateKruskal()
	case 4:
		maze.GenerateBFS()
	default:
		http.Error(w, "Invalid algorithm choice", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maze.Grid)
}

func solveMaze(w http.ResponseWriter, r *http.Request) {
	var req SolveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if maze == nil {
		http.Error(w, "Maze not generated", http.StatusBadRequest)
		return
	}

	switch req.Algo {
	case 1:
		maze.SolveDFS()
	case 2:
		maze.SolveBFS()
	default:
		http.Error(w, "Invalid algorithm choice", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maze)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/generate", generateMaze).Methods("POST")
	r.HandleFunc("/solve", solveMaze).Methods("POST")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	http.Handle("/", corsHandler(r))
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
