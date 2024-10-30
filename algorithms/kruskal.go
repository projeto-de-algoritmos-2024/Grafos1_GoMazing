package algorithms

func (m *Maze) GenerateKruskal() {
	walls := m.getAllWalls()
	sets := makeDisjointSets(m.width, m.height)

	for len(walls) > 0 {
		wall := walls[m.rng.Intn(len(walls))]
		cell1, cell2 := m.getCellsSeparatedByWall(wall)

		if sets.find(cell1) != sets.find(cell2) {
			m.removeWall(cell1, cell2)
			sets.union(cell1, cell2)
		}
		walls = remove(walls, wall)
	}
}
