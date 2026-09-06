import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def convertTemperature(self, celsius: float) -> List[float]:
        return [celsius + 273.15, celsius * 1.8 + 32]


run_tests(
    Solution().convertTemperature,
    [
        {"input": [36.5], "expected": [309.65000, 97.70000]},
        {"input": [122.11], "expected": [395.26000, 251.79800]},
    ],
)
