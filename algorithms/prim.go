package algorithms

import "math/rand"

func (m *Maze) GeneratePrim() {
	walls := make([][2]int, 0, m.width*m.height)
	walls = append(walls, [2]int{0, 0})
	m.Grid[0][0].visited = true

	for len(walls) > 0 {
		currentIndex := m.rng.Intn(len(walls))
		current := walls[currentIndex]
		neighbors := m.getUnvisitedNeighbors(current[0], current[1])

		if len(neighbors) > 0 {
			// Shuffle neighbors to ensure random selection
			rand.Shuffle(len(neighbors), func(i, j int) {
				neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
			})

			next := neighbors[0]
			// Inline removeWall function
			dx, dy := next[0]-current[0], next[1]-current[1]
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

			m.Grid[next[0]][next[1]].visited = true
			walls = append(walls, next)
		} else {
			// Remove the current wall
			walls[currentIndex] = walls[len(walls)-1]
			walls = walls[:len(walls)-1]
		}
	}
}
