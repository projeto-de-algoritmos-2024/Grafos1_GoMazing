package main

import (
	"fmt"

	"github.com/projeto-de-algoritmos-2024/Grafos1_GoMazing/algorithms"
)

func displayMenu() int {
	fmt.Println("Choose an algorithm to generate the maze:")
	fmt.Println("1. Depth-First Search (DFS)")
	fmt.Println("2. Prim's Algorithm")
	fmt.Println("3. Kruskal's Algorithm")
	fmt.Println("4. Breadth-First Search (BFS)")
	fmt.Println("Enter your choice:")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func displaySolveMenu() int {
	fmt.Println("Choose an algorithm to solve the maze:")
	fmt.Println("1. Depth-First Search (DFS)")
	fmt.Println("2. Breadth-First Search (BFS)")
	fmt.Println("Enter your choice:")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func main() {
	choice := displayMenu()
	maze := algorithms.NewMaze(10, 10)

	switch choice {
	case 1:
		maze.GenerateDFS()
	case 2:
		maze.GeneratePrim()
	case 3:
		maze.GenerateKruskal()
	case 4:
		maze.GenerateBFS()
	default:
		fmt.Println("Invalid choice")
		return
	}

	maze.Print()

	solveChoice := displaySolveMenu()

	switch solveChoice {
	case 1:
		maze.SolveDFS()
	case 2:
		maze.SolveBFS()
	default:
		fmt.Println("Invalid choice")
		return
	}

	maze.Print()
}
