import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests

"""
Path existence on a graph built from sorted values.

- The key insight is that nums is sorted, so an edge can only ever matter between neighbouring indices. If
  there's an edge between i and j (i < j), then nums[j] - nums[i] <= maxDiff, and since every gap in between is
  no larger, nums[k+1] - nums[k] <= maxDiff for every k in [i, j) as well. So i and j are connected iff every
  consecutive gap between them is within maxDiff: the graph is just a chain of indices, cut wherever one gap
  exceeds maxDiff, and every connected component is a contiguous run of indices.

- That means we never need to connect i to more than one earlier node. The naive "union i with every earlier
  node within maxDiff" is O(n²) and times out on big inputs where maxDiff is large; linking i to i - 1 only,
  when that gap allows it, is enough.

- worked example, nums = [2, 5, 6, 8], maxDiff = 2:

      index     0        1     2     3
      nums      2   ✗    5     6     8      gaps: 5-2=3 (> 2, cut here), 6-5=1, 8-6=2
      comp      0        1     1     1

  the first gap breaks the chain; everything from index 1 on stays joined. Query [1, 3] compares components
  1 and 1 -> true; query [0, 1] compares 0 and 1 -> false.

Solution - Union-Find:

- Walk the array once and union i with i - 1 whenever nums[i] - nums[i - 1] <= maxDiff. Each query is then a
  "do these two nodes share a root?" check, near O(1) thanks to union-by-rank plus path compression.

- One quirk: it unions on the value nums[i] rather than the index i. That happens to work only because the
  queries also look nodes up by value (nums[q[0]]); unioning on indices would be the cleaner choice.

Solution2 - component labelling:

- Same insight, but we skip Union-Find entirely. Since every component is a contiguous run of indices, a single
  left-to-right sweep can hand each index a component id: start at 0, and bump the id every time a gap exceeds
  maxDiff. Two nodes are connected if they ended up with the same id. O(n) to build, O(1) per query.

Obs: for this problem Solution2 is strictly simpler and faster than the Union-Find version - once we know the
components are contiguous runs, the DSU isn't buying us anything.
"""


class Solution:
    class SparseUnionFind:
        def __init__(self):
            self.parent = {}
            self.rank = {}

        def add(self, x: int):
            if x not in self.parent:
                self.parent[x] = x
                self.rank[x] = 0

        def find(self, x: int) -> int:
            if x not in self.parent:
                self.parent[x] = x
                self.rank[x] = 0
                return x

            if self.parent[x] != x:
                self.parent[x] = self.find(self.parent[x])
            return self.parent[x]

        def union(self, x: int, y: int):
            root_x = self.find(x)
            root_y = self.find(y)

            if root_x == root_y:
                return

            if self.rank[root_x] > self.rank[root_y]:
                self.parent[root_y] = root_x
            elif self.rank[root_x] < self.rank[root_y]:
                self.parent[root_x] = root_y
            else:
                self.parent[root_y] = root_x
                self.rank[root_x] += 1

    def pathExistenceQueries(
        self, n: int, nums: List[int], maxDiff: int, queries: List[List[int]]
    ) -> List[bool]:
        union_find = self.SparseUnionFind()

        for i in range(1, len(nums)):
            if nums[i] - nums[i - 1] <= maxDiff:
                union_find.union(nums[i], nums[i - 1])

        return [
            union_find.find(nums[q[0]]) == union_find.find(nums[q[1]]) for q in queries
        ]


class Solution2:
    def pathExistenceQueries(
        self, n: int, nums: List[int], maxDiff: int, queries: List[List[int]]
    ) -> List[bool]:
        component = [0] * n
        for i in range(1, n):
            component[i] = component[i - 1] + (
                0 if nums[i] - nums[i - 1] <= maxDiff else 1
            )

        return [component[u] == component[v] for u, v in queries]


run_tests(
    Solution().pathExistenceQueries,
    [
        {"input": [2, [1, 3], 1, [[0, 0], [0, 1]]], "expected": [True, False]},
        {
            "input": [4, [2, 5, 6, 8], 2, [[0, 1], [0, 2], [1, 3], [2, 3]]],
            "expected": [False, False, True, True],
        },
    ],
)

run_tests(
    Solution2().pathExistenceQueries,
    [
        {"input": [2, [1, 3], 1, [[0, 0], [0, 1]]], "expected": [True, False]},
        {
            "input": [4, [2, 5, 6, 8], 2, [[0, 1], [0, 2], [1, 3], [2, 3]]],
            "expected": [False, False, True, True],
        },
    ],
)
