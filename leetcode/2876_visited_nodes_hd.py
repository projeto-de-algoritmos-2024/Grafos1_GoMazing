from typing import List


class Solution:
    def countVisitedNodes(self, edges: List[int]) -> List[int]:
        n = len(edges)
        answer = [-1] * n
        visited = [-1] * n

        def dfs(node):
            if visited[node] == 1:
                return answer[node]
            if visited[node] == 0:

                cycle_nodes = []
                current = node
                while True:
                    cycle_nodes.append(current)
                    current = edges[current]
                    if current == node:
                        break
                cycle_length = len(cycle_nodes)
                for n in cycle_nodes:
                    answer[n] = cycle_length
                return cycle_length

            visited[node] = 0

            if edges[node] != -1:
                cycle_length = dfs(edges[node])
                if answer[node] == -1:
                    answer[node] = cycle_length + 1 if cycle_length > 0 else -1
            else:
                answer[node] = 1

            visited[node] = 1
            return answer[node]

        for i in range(n):
            if visited[i] == -1:
                dfs(i)

        return answer
