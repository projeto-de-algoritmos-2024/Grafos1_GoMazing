package main

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	var first, second, last []int
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if parent[v] != v {
			first = []int{parent[v], v}
			second = []int{u, v}
		} else {
			parent[v] = u
		}
	}

	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	for _, edge := range edges {
		if len(second) > 0 && edge[0] == second[0] && edge[1] == second[1] {
			continue
		}
		if !union(edge[0], edge[1], parent) {
			last = edge
		}
	}

	if len(second) == 0 {
		return last
	}
	if len(last) == 0 {
		return second
	}
	return first
}

func find(x int, parent []int) int {
	if parent[x] != x {
		parent[x] = find(parent[x], parent)
	}
	return parent[x]
}

func union(x, y int, parent []int) bool {
	rootX, rootY := find(x, parent), find(y, parent)
	if rootX == rootY {
		return false
	}
	parent[rootX] = rootY
	return true
}
