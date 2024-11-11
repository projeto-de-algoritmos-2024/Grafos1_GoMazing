package main

import (
	"math"
)

func findRotateSteps(ring string, key string) int {
	n, m := len(ring), len(key)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[m] = make([]int, n)

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if ring[j] == key[i] {
				for k := 0; k < n; k++ {
					if ring[k] == key[i+1] {
						diff := int(math.Abs(float64(j - k)))
						step := min(diff, n-diff)
						dp[i][j] = min(dp[i][j], dp[i+1][k]+step)
					}
				}
			}
		}
	}

	result := math.MaxInt32
	for j := 0; j < n; j++ {
		if ring[j] == key[0] {
			diff := int(math.Abs(float64(j - 0)))
			step := min(diff, n-diff)
			result = min(result, dp[0][j]+step)
		}
	}

	return result + m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
