package algorithms

func (m *Maze) GenerateDFS() {
	stack := [][2]int{{0, 0}}
	m.Grid[0][0].visited = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		neighbors := m.getUnvisitedNeighbors(current[0], current[1])

		if len(neighbors) > 0 {
			next := neighbors[m.rng.Intn(len(neighbors))]
			m.removeWall(current, next)
			stack = append(stack, next)
			m.Grid[next[0]][next[1]].visited = true
		} else {
			stack = stack[:len(stack)-1]
		}
	}
}
