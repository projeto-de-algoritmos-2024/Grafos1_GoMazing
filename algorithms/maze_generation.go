package algorithms

import (
	"math/rand"
)

func (m *Maze) MazeGenerateDFS() {
	stack := []*Cell{m.Current}
	m.Current.visited = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		next := m.checkNeighbors(current)
		if next != nil {
			next.visited = true
			stack = append(stack, next)
			m.removeWalls(current, next)
		} else {
			stack = stack[:len(stack)-1]
		}
		m.Steps = append(m.Steps, m.copyGrid())
	}
}

func (m *Maze) MazeGenerateBFS() {
	queue := []*Cell{m.Current}
	m.Current.visited = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		next := m.checkNeighbors(current)
		if next != nil {
			next.visited = true
			queue = append(queue, next)
			m.removeWalls(current, next)
		}
		m.Steps = append(m.Steps, m.copyGrid())
	}
}

func (m *Maze) MazeGeneratePrim() {
	walls := [][2]int{{0, 0}}
	m.Grid[0][0].visited = true

	for len(walls) > 0 {
		current := walls[m.rng.Intn(len(walls))]
		neighbors := m.getUnvisitedNeighbors(current[0], current[1])

		if len(neighbors) > 0 {
			next := neighbors[m.rng.Intn(len(neighbors))]
			m.removeWall(current, next)
			m.Grid[next[0]][next[1]].visited = true
			walls = append(walls, next)
		} else {
			walls = remove(walls, current)
		}
		m.Steps = append(m.Steps, m.copyGrid())
	}
}

func (m *Maze) checkNeighbors(cell *Cell) *Cell {
	neighbors := []*Cell{}

	if top := m.getCell(cell.X, cell.Y-1); top != nil && !top.visited {
		neighbors = append(neighbors, top)
	}
	if right := m.getCell(cell.X+1, cell.Y); right != nil && !right.visited {
		neighbors = append(neighbors, right)
	}
	if bottom := m.getCell(cell.X, cell.Y+1); bottom != nil && !bottom.visited {
		neighbors = append(neighbors, bottom)
	}
	if left := m.getCell(cell.X-1, cell.Y); left != nil && !left.visited {
		neighbors = append(neighbors, left)
	}

	if len(neighbors) > 0 {
		return neighbors[rand.Intn(len(neighbors))]
	}
	return nil
}

func (m *Maze) getCell(x, y int) *Cell {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return nil
	}
	return &m.Grid[y][x]
}

func (m *Maze) removeWalls(current, next *Cell) {
	dx := current.X - next.X
	if dx == 1 {
		current.Walls[3] = false
		next.Walls[1] = false
	} else if dx == -1 {
		current.Walls[1] = false
		next.Walls[3] = false
	}
	dy := current.Y - next.Y
	if dy == 1 {
		current.Walls[0] = false
		next.Walls[2] = false
	} else if dy == -1 {
		current.Walls[2] = false
		next.Walls[0] = false
	}
}
