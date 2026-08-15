import sys
from pathlib import Path
from typing import List
from collections import defaultdict
from functools import cmp_to_key

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
        freqs = defaultdict(int)

        for i in range(k):
            freqs[nums[i]] += 1

        result = []

        sorted_freqs = sorted(freqs.items(), key=cmp_to_key(self.compare_frequencies))
        x_sum = sum(freq[0] * freq[1] for freq in sorted_freqs[:x])
        result.append(x_sum)

        for i in range(1, len(nums) - k + 1):
            freqs[nums[i - 1]] -= 1
            if freqs[nums[i - 1]] == 0:
                del freqs[nums[i - 1]]
            freqs[nums[i + k - 1]] += 1

            sorted_freqs = sorted(
                freqs.items(), key=cmp_to_key(self.compare_frequencies)
            )
            x_sum = sum(freq[0] * freq[1] for freq in sorted_freqs[:x])
            result.append(x_sum)

        return result

    def compare_frequencies(self, f1, f2):
        if f1[1] < f2[1]:
            return 1
        elif f1[1] > f2[1]:
            return -1
        else:
            if f1[0] < f2[0]:
                return 1
            return -1


run_tests(Solution().findXSum, [
    {"input": [[1, 1, 2, 2, 3, 4, 2, 3], 6, 2], "expected": [6, 10, 12]},
    {"input": [[3, 8, 7, 8, 7, 5], 2, 2], "expected": [11, 15, 15, 15, 12]},
    {"input": [[9, 2, 2], 3, 3], "expected": [13]},
])
