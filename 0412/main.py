import sys
from pathlib import Path
from typing import List

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        answer = ["" for _ in range(n)]
        countThree = 0
        countFive = 0

        for i in range(n):
            countThree += 1
            countFive += 1

            if countThree == 3 and countFive == 5:
                answer[i] = "FizzBuzz"
                countThree = 0
                countFive = 0
                continue
            elif countThree == 3:
                answer[i] = "Fizz"
                countThree = 0
                continue
            elif countFive == 5:
                answer[i] = "Buzz"
                countFive = 0
                continue
            else:
                answer[i] = str(i+1)

        return answer


run_tests(Solution().fizzBuzz, [
    {"input": [3], "expected": ["1", "2", "Fizz"]},
    {"input": [5], "expected": ["1", "2", "Fizz", "4", "Buzz"]},
    {
        "input": [15],
        "expected": [
            "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
            "11", "Fizz", "13", "14", "FizzBuzz",
        ],
    },
])