package algorithms

func (m *Maze) GenerateBFS() {
	queue := [][2]int{{0, 0}}
	m.Grid[0][0].visited = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		neighbors := m.getUnvisitedNeighbors(current[0], current[1])

		if len(neighbors) > 0 {
			next := neighbors[m.rng.Intn(len(neighbors))]
			m.removeWall(current, next)
			queue = append(queue, next)
			m.Grid[next[0]][next[1]].visited = true
		}
	}
}
