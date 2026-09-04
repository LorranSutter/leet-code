import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        # Keep track of the letter being compared
        dup = s[0]
        # Count how many duplicated letters were seen
        count = k - 1

        i = 1
        while i < len(s):
            if dup == s[i]:
                count -= 1
                if count == 0:  # Found a sequence of k equal letters
                    s = s[: i - k + 1] + s[i + 1 :]  # Remove duplicated
                    if s == "":  # Case where new string becomes empty
                        break

                    # Move pointer back k positions before the first duplicated letter
                    # We could go before 0, so we do max(0,new_i)
                    i = max(0, i - 2 * k)
                    dup = s[i]
                    count = k - 1
            else:
                # Didn't find duplicate, restart
                dup = s[i]
                count = k - 1
            i += 1

        return s


run_tests(
    Solution().removeDuplicates,
    [
        {"input": ["abcd", 2], "expected": "abcd"},
        {"input": ["deeedbbcccbdaa", 3], "expected": "aa"},
        {"input": ["pbbcggttciiippooaais", 2], "expected": "ps"},
        {
            "input": [
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbb",
                30,
            ],
            "expected": "aaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbb",
        },
        {
            "input": ["yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy", 4],
            "expected": "ybth",
        },
        {"input": ["aaaa", 2], "expected": ""},
    ],
)
