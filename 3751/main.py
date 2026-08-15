import math
import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def totalWaviness(self, num1: int, num2: int) -> int:
        waviness_count = 0
        for i in range(num1, num2 + 1):
            length = math.ceil(math.log10(i))
            if length < 3:
                continue

            digits = [0 for _ in range(length)]
            for j in range(length):
                digits[j] = i % 10
                i //= 10
            
            for j in range(1, length - 1):
                if ((digits[j] > digits[j - 1] and digits[j] > digits[j + 1]) or
                        (digits[j] < digits[j - 1] and digits[j] < digits[j + 1])):
                    waviness_count += 1
        
        return waviness_count


run_tests(Solution().totalWaviness, [
    {"input": [120, 130], "expected": 3},
    {"input": [198, 202], "expected": 3},
    {"input": [4848, 4848], "expected": 2},
])