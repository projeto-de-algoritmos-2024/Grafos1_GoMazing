package algorithms

import (
	"math/rand"
	"time"
)

func (m *Maze) Generate() {
	rand.Seed(time.Now().UnixNano())
	m.Current.visited = true

	for {
		next := m.checkNeighbors(m.Current)
		if next != nil {
			next.visited = true
			m.Stack = append(m.Stack, m.Current)
			m.removeWalls(m.Current, next)
			m.Current = next
		} else if len(m.Stack) > 0 {
			m.Current = m.Stack[len(m.Stack)-1]
			m.Stack = m.Stack[:len(m.Stack)-1]
		} else {
			break
		}
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

func (m *Maze) ToJSON() [][]Cell {
	return m.Grid
}
