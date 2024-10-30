package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	visited bool
	walls   [4]bool // top, right, bottom, left
}

type Maze struct {
	width, height int
	grid          [][]Cell
}

func NewMaze(width, height int) *Maze {
	maze := &Maze{
		width:  width,
		height: height,
		grid:   make([][]Cell, height),
	}
	for i := range maze.grid {
		maze.grid[i] = make([]Cell, width)
		for j := range maze.grid[i] {
			maze.grid[i][j] = Cell{walls: [4]bool{true, true, true, true}}
		}
	}
	return maze
}

func (m *Maze) GenerateDFS() {
	rand.Seed(time.Now().UnixNano())
	stack := [][2]int{{0, 0}}
	m.grid[0][0].visited = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		neighbors := m.getUnvisitedNeighbors(current[0], current[1])

		if len(neighbors) > 0 {
			next := neighbors[rand.Intn(len(neighbors))]
			m.removeWall(current, next)
			stack = append(stack, next)
			m.grid[next[0]][next[1]].visited = true
		} else {
			stack = stack[:len(stack)-1]
		}
	}
}

func (m *Maze) GeneratePrim() {
	// Implementation of Prim's Algorithm for Maze Generation
}

func (m *Maze) GenerateKruskal() {
	// Implementation of Kruskal's Algorithm for Maze Generation
}

func (m *Maze) GenerateBFS() {
	// Implementation of BFS for Maze Generation
}

func (m *Maze) SolveDFS() {
	// Implementation of DFS for Maze Solving
}

func (m *Maze) SolveBFS() {
	// Implementation of BFS for Maze Solving
}

func (m *Maze) getUnvisitedNeighbors(x, y int) [][2]int {
	neighbors := [][2]int{}
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // top, bottom, left, right

	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && ny >= 0 && nx < m.height && ny < m.width && !m.grid[nx][ny].visited {
			neighbors = append(neighbors, [2]int{nx, ny})
		}
	}
	return neighbors
}

func (m *Maze) removeWall(current, next [2]int) {
	dx, dy := next[0]-current[0], next[1]-current[1]

	if dx == -1 {
		m.grid[current[0]][current[1]].walls[0] = false
		m.grid[next[0]][next[1]].walls[2] = false
	} else if dx == 1 {
		m.grid[current[0]][current[1]].walls[2] = false
		m.grid[next[0]][next[1]].walls[0] = false
	} else if dy == -1 {
		m.grid[current[0]][current[1]].walls[3] = false
		m.grid[next[0]][next[1]].walls[1] = false
	} else if dy == 1 {
		m.grid[current[0]][current[1]].walls[1] = false
		m.grid[next[0]][next[1]].walls[3] = false
	}
}

func (m *Maze) Print() {
	for i := 0; i < m.height; i++ {
		// Print the top walls
		for j := 0; j < m.width; j++ {
			if m.grid[i][j].walls[0] {
				fmt.Print("+---")
			} else {
				fmt.Print("+   ")
			}
		}
		fmt.Println("+")

		// Print the left walls and spaces
		for j := 0; j < m.width; j++ {
			if m.grid[i][j].walls[3] {
				fmt.Print("|   ")
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println("|")
	}

	// Print the bottom walls
	for j := 0; j < m.width; j++ {
		fmt.Print("+---")
	}
	fmt.Println("+")
}

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
	maze := NewMaze(10, 10)

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
