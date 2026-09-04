import sys
from typing import List
from collections import defaultdict
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import equal_matrix_unordered, run_tests


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        result = defaultdict(list)
        for s in strs:
            c = "".join(sorted(s))
            result[c].append(s)

        return list(result.values())


run_tests(
    Solution().groupAnagrams,
    [
        {
            "input": [["eat", "tea", "tan", "ate", "nat", "bat"]],
            "expected": [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
        },
        {
            "input": [[""]],
            "expected": [[""]],
        },
        {
            "input": [["a"]],
            "expected": [["a"]],
        },
    ],
    is_equal=equal_matrix_unordered,
)
