import sys
import heapq
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def kthLargestNumber(self, nums: List[str], k: int) -> str:
        min_heap = []

        for num in nums:
            heapq.heappush(min_heap, int(num))

            if len(min_heap) > k:
                heapq.heappop(min_heap)

        return str(min_heap[0])


run_tests(
    Solution().kthLargestNumber,
    [
        {"input": [["3", "6", "7", "10"], 4], "expected": "3"},
        {"input": [["2", "21", "12", "1"], 3], "expected": "2"},
        {"input": [["0", "0"], 2], "expected": "0"},
    ],
)
