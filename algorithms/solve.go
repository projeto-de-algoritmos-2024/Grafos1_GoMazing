package algorithms

import "fmt"

func (m *Maze) SolveDFS() {
	start := [2]int{0, 0}
	end := [2]int{m.height - 1, m.width - 1}
	stack := [][2]int{start}
	visited := make(map[[2]int]bool)
	visited[start] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current == end {
			fmt.Println("Path found using DFS")
			return
		}

		for _, neighbor := range m.getNeighbors(current[0], current[1]) {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
				visited[neighbor] = true
			}
		}
	}
	fmt.Println("No path found using DFS")
}

func (m *Maze) SolveBFS() {
	start := [2]int{0, 0}
	end := [2]int{m.height - 1, m.width - 1}
	queue := [][2]int{start}
	visited := make(map[[2]int]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			fmt.Println("Path found using BFS")
			return
		}

		for _, neighbor := range m.getNeighbors(current[0], current[1]) {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
	fmt.Println("No path found using BFS")
}
