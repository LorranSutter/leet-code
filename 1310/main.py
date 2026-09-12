import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def xorQueries(self, arr: List[int], queries: List[List[int]]) -> List[int]:
        prefix_xor = [0] * (len(arr) + 1)
        for i in range(len(arr)):
            prefix_xor[i + 1] = prefix_xor[i] ^ arr[i]

        result = [0] * len(queries)
        for i in range(len(queries)):
            result[i] = prefix_xor[queries[i][1] + 1] ^ prefix_xor[queries[i][0]]

        return result


run_tests(
    Solution().xorQueries,
    [
        {
            "input": [[1, 3, 4, 8], [[0, 1], [1, 2], [0, 3], [3, 3]]],
            "expected": [2, 7, 14, 8],
        },
        {
            "input": [[4, 8, 2, 10], [[2, 3], [1, 3], [0, 0], [0, 3]]],
            "expected": [8, 0, 4, 4],
        },
        {"input": [[16], [[0, 0], [0, 0], [0, 0]]], "expected": [16, 16, 16]},
    ],
)
