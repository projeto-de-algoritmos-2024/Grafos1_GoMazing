package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

type Maze struct {
	width, height int
	grid          [][]Cell
	rng           *rand.Rand
	Steps         [][]Cell // Change this to [][]Cell
}

type Cell struct {
	visited bool
	walls   [4]bool // top, right, bottom, left
}

type DisjointSets struct {
	parent, rank []int
}

func NewMaze(width, height int) *Maze {
	maze := &Maze{
		width:  width,
		height: height,
		grid:   make([][]Cell, height),
		rng:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Steps:  make([][]Cell, 0), // Initialize the steps
	}
	for i := range maze.grid {
		maze.grid[i] = make([]Cell, width)
		for j := range maze.grid[i] {
			maze.grid[i][j] = Cell{walls: [4]bool{true, true, true, true}}
		}
	}
	return maze
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

	gridCopy := m.copyGrid()
	m.Steps = append(m.Steps, gridCopy...)
}

func (m *Maze) getAllWalls() [][2]int {
	// Implementation to get all walls
	return nil
}

func (m *Maze) getCellsSeparatedByWall(wall [2]int) ([2]int, [2]int) {
	// Implementation to get cells separated by a wall
	return [2]int{}, [2]int{}
}

func makeDisjointSets(width, height int) *DisjointSets {
	size := width * height
	ds := &DisjointSets{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := range ds.parent {
		ds.parent[i] = i
	}
	return ds
}

func (ds *DisjointSets) find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.find(ds.parent[x])
	}
	return ds.parent[x]
}

func (ds *DisjointSets) union(x, y int) {
	rootX := ds.find(x)
	rootY := ds.find(y)
	if rootX != rootY {
		if ds.rank[rootX] > ds.rank[rootY] {
			ds.parent[rootY] = rootX
		} else if ds.rank[rootX] < ds.rank[rootY] {
			ds.parent[rootX] = rootY
		} else {
			ds.parent[rootY] = rootX
			ds.rank[rootX]++
		}
	}
}

func remove(slice [][2]int, item [2]int) [][2]int {
	for i, other := range slice {
		if other == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func (m *Maze) getNeighbors(x, y int) [][2]int {
	neighbors := [][2]int{}
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // top, bottom, left, right

	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && ny >= 0 && nx < m.height && ny < m.width {
			neighbors = append(neighbors, [2]int{nx, ny})
		}
	}
	return neighbors
}

func (m *Maze) coordToIndex(x, y int) int {
	return x*m.width + y
}

func (m *Maze) indexToCoord(index int) (int, int) {
	return index / m.width, index % m.width
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

func (m *Maze) copyGrid() [][]Cell {
	copy := make([][]Cell, len(m.grid))
	for i := range m.grid {
		copy[i] = make([]Cell, len(m.grid[i]))
		for j := range m.grid[i] {
			copy[i][j] = m.grid[i][j]
		}
	}
	return copy
}
