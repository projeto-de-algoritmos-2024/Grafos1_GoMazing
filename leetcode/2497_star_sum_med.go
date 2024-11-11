package main

import (
	"fmt"
	"sort"
)

func maxStarSum(vals []int, edges [][]int, k int) int {
	n := len(vals)
	adj := make([][]int, n)

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	maxSum := vals[0]

	for i := 0; i < n; i++ {
		neighbors := adj[i]
		sort.Slice(neighbors, func(a, b int) bool {
			return vals[neighbors[a]] > vals[neighbors[b]]
		})

		starSum := vals[i]
		for j := 0; j < k && j < len(neighbors); j++ {
			if vals[neighbors[j]] > 0 {
				starSum += vals[neighbors[j]]
			}
		}

		if starSum > maxSum {
			maxSum = starSum
		}
	}

	return maxSum
}

func mainLeet() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	vals := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vals[i])
	}

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = make([]int, 2)
		fmt.Scan(&edges[i][0], &edges[i][1])
	}

	fmt.Println(maxStarSum(vals, edges, k))
}

func main() {
	mainLeet()
}
