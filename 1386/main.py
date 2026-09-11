import sys
from typing import List
from pathlib import Path
from collections import defaultdict

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests

"""
Preprocessing:
- We create a map from row number to the set of reserved seats in that row. This makes the iteration
  simpler, since we just have to check a number instead of a tuple, and any row that never shows up in
  reservedSeats needs no checking at all.

- The key insight is that seats 1 and 10 belong to none of the three blocks, so they can't matter and we
  ignore them. That leaves three blocks over seats 2-9, and here's how they line up:

          left   : 2 3 4 5
          middle :     4 5 6 7
          right  :         6 7 8 9

- Left and right don't share a seat, so a completely free row always fits exactly 2 groups. Middle shares
  seats 4-5 with left and 6-7 with right, so choosing middle rules both of the others out.

- Rows with no reservations are handled in bulk: we start the count at 2 * (n - rows in the map), then only
  iterate over the rows that actually have a reservation.

- For each of those rows there are four situations, which the if/elif chain walks in this order:
    1. left is free                 -> +1 group
    2. left and right both free     -> +2 groups (the nested check right after situation 1)
    3. right is free but not left   -> +1 group
    4. only the middle is free      -> +1 group

- Situation 4 is really just a fallback: we only reach it once both left and right are known to be blocked,
  and then seats 4-5-6-7 might still be open for a single group. For example, a row with seats 2 and 3
  reserved kills the left block, but 4-5-6-7 are untouched, so we still seat one group down the middle.
"""


class Solution:
    def maxNumberOfFamilies(self, n: int, reservedSeats: List[List[int]]) -> int:
        reservedSeatsMap = defaultdict(set)
        for s in reservedSeats:
            reservedSeatsMap[s[0] - 1].add(s[1])
        # {1: {2,3,8}, 2: {6}, 3: {1,10}}

        left = [2, 3, 4, 5]
        middle = [4, 5, 6, 7]
        right = [6, 7, 8, 9]

        groups = 2 * (n - len(reservedSeatsMap))
        for i in reservedSeatsMap.keys():
            if not any(s in reservedSeatsMap[i] for s in left):
                groups += 1
                if not any(s in reservedSeatsMap[i] for s in right):
                    groups += 1
            elif not any(s in reservedSeatsMap[i] for s in right):
                groups += 1
            elif not any(s in reservedSeatsMap[i] for s in middle):
                groups += 1

        return groups


run_tests(
    Solution().solve,
    [
        {
            "input": [3, [[1, 2], [1, 3], [1, 8], [2, 6], [3, 1], [3, 10]]],
            "expected": 4,
        },
        {"input": [2, [[2, 1], [1, 8], [2, 6]]], "expected": 2},
        {"input": [4, [[4, 3], [1, 4], [4, 6], [1, 7]]], "expected": 4},
        {"input": [3, [[2, 3]]], "expected": 5},
    ],
)
