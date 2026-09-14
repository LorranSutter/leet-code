import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def countSpecialIntegers(self, nums: list[int]) -> int:
        seen = set()
        specials = set()
        current = 0
        for num in nums:
            if num in specials:
                if num != current:
                    current = num
                    specials.remove(num)
                    seen.add(num)
            else:
                if num not in seen:
                    specials.add(num)
                current = num
        return len(specials)


run_tests(
    Solution().countSpecialIntegers,
    [
        {"input": [[1, 2, 2, 1]], "expected": 1},
        {"input": [[3, 3, 1, 2, 2, 1]], "expected": 2},
    ],
)
