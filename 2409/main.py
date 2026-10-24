import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def countDaysTogether(
        self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str
    ) -> int:
        days_per_month = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

        arriveAlice = list(map(int, arriveAlice.split("-")))
        leaveAlice = list(map(int, leaveAlice.split("-")))
        arriveBob = list(map(int, arriveBob.split("-")))
        leaveBob = list(map(int, leaveBob.split("-")))

        alice = [0, 0]
        alice[0] = sum(days_per_month[: arriveAlice[0] - 1]) + arriveAlice[1]
        alice[1] = sum(days_per_month[: leaveAlice[0] - 1]) + leaveAlice[1]

        bob = [0, 0]
        bob[0] = sum(days_per_month[: arriveBob[0] - 1]) + arriveBob[1]
        bob[1] = sum(days_per_month[: leaveBob[0] - 1]) + leaveBob[1]

        overlap = min(alice[1], bob[1]) - max(alice[0], bob[0])
        if overlap >= 0:
            return overlap + 1
        return 0


run_tests(
    Solution().countDaysTogether,
    [
        {"input": ["08-15", "08-18", "08-16", "08-19"], "expected": 3},
        {"input": ["10-01", "10-31", "11-01", "12-31"], "expected": 0},
        {"input": ["09-01", "10-19", "06-19", "10-20"], "expected": 49},
        {"input": ["08-15", "08-18", "08-16", "08-16"], "expected": 1},
    ],
)
