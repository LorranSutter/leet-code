import sys
from typing import List
from pathlib import Path
from collections import Counter

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counter = Counter(nums)
        most_frequent = counter.most_common(k)

        return [item[0] for item in most_frequent]


run_tests(
    Solution().topKFrequent,
    [
        {"input": [[1, 1, 1, 2, 2, 3], 2], "expected": [1, 2]},
        {"input": [[1], 1], "expected": [1]},
        {"input": [[1, 2, 1, 2, 1, 2, 3, 1, 3, 2], 2], "expected": [1, 2]},
    ],
)
