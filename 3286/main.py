import sys
from typing import List
from pathlib import Path
from collections import deque

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests

"""
Solution:

- The problem reduces to a shortest-path search: instead of minimizing distance, we minimize the number of unsafe
  cells (the 1s) that the walk steps on, since each one costs exactly one health point.
- Dijkstra would solve this too, but it's overkill here - its O(n log n) priority queue is wasted when every edge
  weight is only 0 or 1.
- That's exactly the case 0-1 BFS is built for: a specialized BFS for graphs whose edge weights are only 0 or 1.
- Instead of a priority queue, it uses a deque that can be pushed to from either end. A move onto a safe cell
  (weight 0) is pushed to the front, and a move onto an unsafe cell (weight 1) is pushed to the back. Since 0-cost
  moves always jump the queue ahead of 1-cost ones, the deque stays sorted by distance without ever needing to sort
  it explicitly.
- One subtlety: `dist[(0, 0)]` is seeded with `grid[0][0]` rather than 0, because the starting cell itself can be
  unsafe and cost a health point. The first cell popped whose position equals the target is guaranteed to hold the
  minimum unsafe-cell count, since 0-1 BFS always pops nodes in non-decreasing distance order - so we return as
  soon as we reach it.
- `findSafeWalk` then just checks whether `health` has at least one point left over after paying for that minimum
  count: `health - bfs_01() >= 1`.
"""


class Solution:
    def findSafeWalk(self, grid: List[List[int]], health: int) -> bool:
        m, n = len(grid), len(grid[0])
        directions = ((-1, 0), (1, 0), (0, 1), (0, -1))

        def bfs_01() -> int:
            dist = {}
            for i in range(m):
                for j in range(n):
                    dist[(i, j)] = float("inf")

            queue = deque()
            queue.append((0, 0))
            # Populate the first dist, bc the first cell could be 1
            dist[(0, 0)] = grid[0][0]

            while queue:
                u = queue.popleft()

                if u == (m - 1, n - 1):
                    return dist[u]

                for di, dj in directions:
                    dx, dy = u[0] + di, u[1] + dj
                    if dx < 0 or dy < 0 or dx >= m or dy >= n:
                        continue

                    v = (dx, dy)
                    weight = grid[dx][dy]

                    if dist[u] + weight < dist[v]:
                        dist[v] = dist[u] + weight

                        if weight == 0:
                            queue.appendleft(v)
                        else:
                            queue.append(v)

            return dist[(m - 1, n - 1)]

        return health - bfs_01() >= 1


run_tests(
    Solution().findSafeWalk,
    [
        {
            "input": [[[0, 1, 0, 0, 0], [0, 1, 0, 1, 0], [0, 0, 0, 1, 0]], 1],
            "expected": True,
        },
        {
            "input": [
                [
                    [0, 1, 1, 0, 0, 0],
                    [1, 0, 1, 0, 0, 0],
                    [0, 1, 1, 1, 0, 1],
                    [0, 0, 1, 0, 1, 0],
                ],
                3,
            ],
            "expected": False,
        },
        {"input": [[[1, 1, 1], [1, 0, 1], [1, 1, 1]], 5], "expected": True},
    ],
)
