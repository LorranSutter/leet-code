import sys
import heapq
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        min_heap = []

        for num in nums:
            heapq.heappush(min_heap, num)

            if len(min_heap) > k:
                heapq.heappop(min_heap)

        return min_heap[0]


run_tests(
    Solution().findKthLargest,
    [
        {"input": [[3, 2, 1, 5, 6, 4], 2], "expected": 5},
        {"input": [[3, 2, 3, 1, 2, 4, 5, 5, 6], 4], "expected": 4},
    ],
)
