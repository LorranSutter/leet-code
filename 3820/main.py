import sys
from typing import Dict, List
from pathlib import Path
from collections import defaultdict, deque

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def specialNodes(
        self, n: int, edges: List[List[int]], x: int, y: int, z: int
    ) -> int:
        graph = defaultdict(set)

        for e in edges:
            graph[e[0]].add(e[1])
            graph[e[1]].add(e[0])

        def calculate_distances(start: int) -> Dict[int, int]:
            queue = deque()
            queue.append((start, 0))
            dists = {}
            visited = set()

            while queue:
                u, d = queue.popleft()

                dists[u] = d
                visited.add(u)

                for v in graph[u]:
                    if v not in visited:
                        queue.append((v, d + 1))

            return dists

        x_dists = calculate_distances(x)
        y_dists = calculate_distances(y)
        z_dists = calculate_distances(z)

        triplets = 0
        for node in graph.keys():
            d = sorted([x_dists[node], y_dists[node], z_dists[node]])

            if d[0] ** 2 + d[1] ** 2 == d[2] ** 2:
                triplets += 1

        return triplets


run_tests(
    Solution().specialNodes,
    [
        {"input": [4, [[0, 1], [0, 2], [0, 3]], 1, 2, 3], "expected": 3},
        {"input": [4, [[0, 1], [1, 2], [2, 3]], 0, 3, 2], "expected": 0},
        {"input": [4, [[0, 1], [1, 2], [1, 3]], 1, 3, 0], "expected": 1},
    ],
)
