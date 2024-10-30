package algorithms

func (m *Maze) GeneratePrim() {
	walls := [][2]int{{0, 0}}
	m.grid[0][0].visited = true

	for len(walls) > 0 {
		current := walls[m.rng.Intn(len(walls))]
		neighbors := m.getUnvisitedNeighbors(current[0], current[1])

		if len(neighbors) > 0 {
			next := neighbors[m.rng.Intn(len(neighbors))]
			m.removeWall(current, next)
			m.grid[next[0]][next[1]].visited = true
			walls = append(walls, next)
		} else {
			walls = remove(walls, current)
		}
	}
}
