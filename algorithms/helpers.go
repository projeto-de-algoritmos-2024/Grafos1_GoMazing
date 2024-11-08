package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

type Maze struct {
	width, height int
	Grid          [][]Cell
	rng           *rand.Rand
	Steps         [][][]Cell
	Current       *Cell
	Stack         []*Cell
}

type Cell struct {
	X, Y    int
	visited bool
	Walls   [4]bool // top, right, bottom, left
}

type DisjointSets struct {
	parent, rank []int
}

func NewMaze(width, height int) *Maze {
	maze := &Maze{
		width:  width,
		height: height,
		Grid:   make([][]Cell, height),
		rng:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Steps:  make([][][]Cell, 0),
	}
	for i := range maze.Grid {
		maze.Grid[i] = make([]Cell, width)
		for j := range maze.Grid[i] {
			maze.Grid[i][j] = Cell{X: j, Y: i, Walls: [4]bool{true, true, true, true}}
		}
	}
	maze.Current = &maze.Grid[0][0]
	maze.Stack = []*Cell{}
	return maze
}

func (m *Maze) getAllWalls() [][2][2]int {
	var walls [][2][2]int
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if x < m.width-1 {
				walls = append(walls, [2][2]int{{x, y}, {x + 1, y}})
			}
			if y < m.height-1 {
				walls = append(walls, [2][2]int{{x, y}, {x, y + 1}})
			}
		}
	}
	return walls
}

func (m *Maze) getCellsSeparatedByWall(wall [2][2]int) ([2]int, [2]int) {
	return wall[0], wall[1]
}

func (m *Maze) GenerateMaze(x, y int) {
	m.Grid[x][y].visited = true
	m.Steps = append(m.Steps, m.copyGrid()) // Record each step for visualization

	for {
		neighbors := m.getUnvisitedNeighbors(x, y)
		if len(neighbors) == 0 {
			break // No unvisited neighbors, backtracking stops
		}

		// Randomly pick an unvisited neighbor
		next := neighbors[m.rng.Intn(len(neighbors))]
		m.removeWall([2]int{x, y}, next)
		m.GenerateMaze(next[0], next[1]) // Recursively visit next cell
	}
}

func (m *Maze) getUnvisitedNeighbors(x, y int) [][2]int {
	neighbors := [][2]int{}
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // top, bottom, left, right

	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && ny >= 0 && nx < m.height && ny < m.width && !m.Grid[nx][ny].visited {
			neighbors = append(neighbors, [2]int{nx, ny})
		}
	}
	return neighbors
}

func (m *Maze) removeWall(current, next [2]int) {
	dx, dy := next[0]-current[0], next[1]-current[1]

	// Adjust the walls based on position
	if dx == -1 {
		m.Grid[current[0]][current[1]].Walls[0] = false
		m.Grid[next[0]][next[1]].Walls[2] = false
	} else if dx == 1 {
		m.Grid[current[0]][current[1]].Walls[2] = false
		m.Grid[next[0]][next[1]].Walls[0] = false
	} else if dy == -1 {
		m.Grid[current[0]][current[1]].Walls[3] = false
		m.Grid[next[0]][next[1]].Walls[1] = false
	} else if dy == 1 {
		m.Grid[current[0]][current[1]].Walls[1] = false
		m.Grid[next[0]][next[1]].Walls[3] = false
	}
	m.Steps = append(m.Steps, m.copyGrid()) // Save each step
}

func (m *Maze) copyGrid() [][]Cell {
	copy := make([][]Cell, len(m.Grid))
	for i := range m.Grid {
		copy[i] = make([]Cell, len(m.Grid[i]))
		for j := range m.Grid[i] {
			copy[i][j] = m.Grid[i][j]
		}
	}
	return copy
}

func (m *Maze) Print() {
	for i := 0; i < m.height; i++ {
		// Print the top walls
		for j := 0; j < m.width; j++ {
			if m.Grid[i][j].Walls[0] {
				fmt.Print("+---")
			} else {
				fmt.Print("+   ")
			}
		}
		fmt.Println("+")

		// Print the left walls and spaces
		for j := 0; j < m.width; j++ {
			if m.Grid[i][j].Walls[3] {
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
