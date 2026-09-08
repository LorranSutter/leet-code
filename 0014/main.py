import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        prefix = ""

        valid = True
        index = 0
        while True:
            if index >= len(strs[0]):
                break

            letter = strs[0][index]
            for i in range(len(strs)):
                if index >= len(strs[i]):
                    valid = False
                    break
                if letter != strs[i][index]:
                    valid = False
                    break

            if not valid:
                break

            prefix += letter
            index += 1

        return prefix


run_tests(
    Solution().longestCommonPrefix,
    [
        {"input": [["flower", "flow", "flight"]], "expected": "fl"},
        {"input": [["dog", "racecar", "car"]], "expected": ""},
        {"input": [[""]], "expected": ""},
    ],
)
