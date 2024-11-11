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

	find := func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		rootX, rootY := find(x), find(y)
		if rootX == rootY {
			return false
		}
		parent[rootX] = rootY
		return true
	}

	for _, edge := range edges {
		if len(second) > 0 && edge[0] == second[0] && edge[1] == second[1] {
			continue
		}
		if !union(edge[0], edge[1]) {
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
