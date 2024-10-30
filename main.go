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

func main() {
	maze := NewMaze(10, 10)
	maze.GenerateDFS()
	maze.Print()
}
