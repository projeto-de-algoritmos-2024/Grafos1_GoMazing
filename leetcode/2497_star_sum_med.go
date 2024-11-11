package main

import (
	"sort"
)

// https://leetcode.com/problems/maximum-star-sum-of-a-graph/
func maxStarSum(vals []int, edges [][]int, k int) int {
	n := len(vals)

	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	maxStarSum := vals[0]

	for i := 0; i < n; i++ {

		var neighborVals []int
		for _, neighbor := range graph[i] {
			neighborVals = append(neighborVals, vals[neighbor])
		}

		sort.Slice(neighborVals, func(a, b int) bool {
			return neighborVals[a] > neighborVals[b]
		})

		currentSum := vals[i]

		for j := 0; j < k && j < len(neighborVals); j++ {
			if neighborVals[j] > 0 {
				currentSum += neighborVals[j]
			}
		}

		if currentSum > maxStarSum {
			maxStarSum = currentSum
		}
	}

	return maxStarSum
}
