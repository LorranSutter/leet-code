import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def generateTag(self, caption: str) -> str:
        result = ""
        for c in caption.split():
            result += c.lower().capitalize()

        if result == "":
            return "#"

        result = "#" + result[0].lower() + result[1:]
        return result[:100]


run_tests(
    Solution().generateTag,
    [
        {
            "input": ["Leetcode daily streak achieved"],
            "expected": "#leetcodeDailyStreakAchieved",
        },
        {"input": ["can I Go There"], "expected": "#canIGoThere"},
        {
            "input": [
                "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"
            ],
            "expected": "#hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh",
        },
        {
            "input": [" Leetcode    daily    streak     achieved"],
            "expected": "#leetcodeDailyStreakAchieved",
        },
        {"input": ["   "], "expected": "#"},
    ],
)
