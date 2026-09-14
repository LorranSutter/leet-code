import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    class UnionFind:
        def __init__(self):
            self.parent = {}
            self.rank = {}

        def find(self, u: int) -> int:
            if u not in self.parent:
                self.parent[u] = u
                self.rank[u] = 0
                return u

            if self.parent[u] != u:
                self.parent[u] = self.find(self.parent[u])
            return self.parent[u]

        def union(self, u: int, v: int):
            root_u = self.find(u)
            root_v = self.find(v)

            if root_u == root_v:
                return

            if self.rank[root_u] > self.rank[root_v]:
                self.parent[root_v] = root_u
            elif self.rank[root_u] < self.rank[root_v]:
                self.parent[root_u] = root_v
            else:
                self.parent[root_v] = root_u
                self.rank[root_u] += 1

    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        union_find = self.UnionFind()

        remove_candidates = set()
        for edge in edges:
            if union_find.find(edge[0]) == union_find.find(edge[1]):
                remove_candidates.add(tuple(edge))
            else:
                union_find.union(edge[0], edge[1])

        if len(remove_candidates) > 1:
            for edge in edges[::-1]:
                if tuple(edge) in remove_candidates:
                    return edge
        elif len(remove_candidates) == 1:
            return list(list(remove_candidates)[0])

        return []


run_tests(
    Solution().findRedundantConnection,
    [
        {"input": [[[1, 2], [1, 3], [2, 3]]], "expected": [2, 3]},
        {"input": [[[1, 2], [2, 3], [3, 4], [1, 4], [1, 5]]], "expected": [1, 4]},
        {
            "input": [
                [
                    [3, 7],
                    [1, 4],
                    [2, 8],
                    [1, 6],
                    [7, 9],
                    [6, 10],
                    [1, 7],
                    [2, 3],
                    [8, 9],
                    [5, 9],
                ]
            ],
            "expected": [8, 9],
        },
    ],
)
