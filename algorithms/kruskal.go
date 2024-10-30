package algorithms

func (m *Maze) GenerateKruskal() {
	walls := m.getAllWalls()
	sets := makeDisjointSets(m.width, m.height)

	for len(walls) > 0 {
		wall := walls[m.rng.Intn(len(walls))]
		cell1, cell2 := m.getCellsSeparatedByWall(wall)

		index1 := m.coordToIndex(cell1[0], cell1[1])
		index2 := m.coordToIndex(cell2[0], cell2[1])

		if sets.find(index1) != sets.find(index2) {
			m.removeWall(cell1, cell2)
			sets.union(index1, index2)
		}
		walls = remove(walls, wall)
	}
}
