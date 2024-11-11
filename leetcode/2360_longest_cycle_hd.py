from typing import List


class Solution:
    def longestCycle(self, edges: List[int]) -> int:
        n = len(edges)
        visited = [-1] * n
        longest = -1

        def dfs(node, depth, path):
            nonlocal longest
            if visited[node] != -1:
                if node in path:
                    cycle_length = depth - path[node]
                    longest = max(longest, cycle_length)
                return
            visited[node] = depth
            path[node] = depth
            if edges[node] != -1:
                dfs(edges[node], depth + 1, path)
            path.pop(node)

        for i in range(n):
            if visited[i] == -1:
                dfs(i, 0, {})

        return longest
