package algorithms

import "math/rand"

func (m *Maze) GenerateDFS() {
	stack := make([][2]int, 0, m.width*m.height)
	stack = append(stack, [2]int{0, 0})
	m.Grid[0][0].visited = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
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

			stack = append(stack, next)
			m.Grid[next[0]][next[1]].visited = true
		} else {
			stack = stack[:len(stack)-1]
		}
	}
}
