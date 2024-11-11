from typing import List


class Solution:
    def longestCycle(self, edges: List[int]) -> int:
        n = len(edges)
        visited = [-1] * n
        longest = -1

        def dfs(node, depth):
            nonlocal longest
            if visited[node] != -1:
                if visited[node] >= depth:
                    longest = max(longest, depth - visited[node])
                return
            visited[node] = depth
            if edges[node] != -1:
                dfs(edges[node], depth + 1)
            visited[node] = -2

        for i in range(n):
            if visited[i] == -1:
                dfs(i, 0)

        return longest
