import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def validIPAddress(self, queryIP: str) -> str:
        def validate_ipv4(ip: List[str]) -> bool:
            for x in ip:
                if 1 < len(x) > 3:
                    return False
                if len(x) > 1 and x[0] == "0":
                    return False
                if not x.isdigit():
                    return False
                if int(x) >= 256:
                    return False

            return True

        def validate_ipv6(ip: List[str]) -> bool:
            for x in ip:
                if 1 < len(x) > 4:
                    return False
                try:
                    int(x, 16)
                except ValueError:
                    return False

            return True

        ipv4_candidate = queryIP.split(".")
        if len(ipv4_candidate) == 4 and validate_ipv4(ipv4_candidate):
            return "IPv4"

        ipv6_candidate = queryIP.split(":")
        if len(ipv6_candidate) == 8 and validate_ipv6(ipv6_candidate):
            return "IPv6"

        return "Neither"


run_tests(
    Solution().validIPAddress,
    [
        {"input": ["172.16.254.1"], "expected": "IPv4"},
        {"input": ["2001:0db8:85a3:0:0:8A2E:0370:7334"], "expected": "IPv6"},
        {"input": ["256.256.256.256"], "expected": "Neither"},
    ],
)
