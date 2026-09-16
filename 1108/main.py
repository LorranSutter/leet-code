import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def defangIPaddr(self, address: str) -> str:
        return address.replace(".", "[.]")


run_tests(
    Solution().defangIPaddr,
    [
        {"input": ["1.1.1.1"], "expected": "1[.]1[.]1[.]1"},
        {"input": ["255.100.50.0"], "expected": "255[.]100[.]50[.]0"},
    ],
)
