import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests

MOD = 10**9 + 7


class Solution:
    def sumAndMultiply(self, s: str, queries: List[List[int]]) -> List[int]:
        n = len(s)
        inv10 = pow(10, MOD - 2, MOD)

        # digit_sum[i] : sum of digits of s[:i]
        # nz[i]        : count of non-zero digits in s[:i]
        # pow10[k]     : 10**k % MOD
        # weighted[i]  : sum over non-zero j < i of digit(j) * inv10**nz[j+1]  (mod MOD)
        digit_sum = [0] * (n + 1)
        nz = [0] * (n + 1)
        pow10 = [1] * (n + 1)
        weighted = [0] * (n + 1)

        inv10_pow = 1  # == inv10 ** nz[i]
        for i in range(1, n + 1):
            d = int(s[i - 1])
            digit_sum[i] = digit_sum[i - 1] + d
            pow10[i] = pow10[i - 1] * 10 % MOD
            nz[i] = nz[i - 1]
            weighted[i] = weighted[i - 1]
            if d != 0:
                nz[i] += 1
                inv10_pow = inv10_pow * inv10 % MOD
                weighted[i] = (weighted[i] + d * inv10_pow) % MOD

        result = []
        for l, r in queries:
            s_sum = digit_sum[r + 1] - digit_sum[l]
            x = pow10[nz[r + 1]] * (weighted[r + 1] - weighted[l]) % MOD
            result.append(x * s_sum % MOD)

        return result


run_tests(
    Solution().sumAndMultiply,
    [
        {"input": ["10203004", [[0, 7], [1, 3], [4, 6]]], "expected": [12340, 4, 9]},
        {"input": ["1000", [[0, 3], [1, 1]]], "expected": [1, 0]},
        {"input": ["9876543210", [[0, 9]]], "expected": [444444137]},
        {
            "input": [
                "57569977386369791",
                [
                    [0, 1],
                    [0, 2],
                    [0, 3],
                    [0, 5],
                    [0, 6],
                    [0, 11],
                    [0, 13],
                    [0, 16],
                    [1, 2],
                    [1, 4],
                    [1, 6],
                    [1, 10],
                    [1, 11],
                    [1, 12],
                    [1, 13],
                    [1, 16],
                    [2, 4],
                    [2, 7],
                    [2, 8],
                    [2, 9],
                    [2, 11],
                    [2, 14],
                    [3, 5],
                    [3, 6],
                    [3, 8],
                    [3, 9],
                    [3, 10],
                    [3, 14],
                    [4, 9],
                    [4, 14],
                    [4, 15],
                    [4, 16],
                    [5, 7],
                    [5, 9],
                    [5, 10],
                    [5, 11],
                    [5, 14],
                    [6, 6],
                    [6, 7],
                    [6, 8],
                    [6, 12],
                    [6, 13],
                    [6, 14],
                    [7, 9],
                    [7, 10],
                    [7, 11],
                    [7, 16],
                    [8, 9],
                    [8, 16],
                    [9, 10],
                    [9, 12],
                    [10, 11],
                    [10, 12],
                    [10, 13],
                    [10, 14],
                    [11, 13],
                    [11, 14],
                    [11, 15],
                    [11, 16],
                    [12, 13],
                    [12, 14],
                    [12, 16],
                    [13, 15],
                    [13, 16],
                    [14, 15],
                    [14, 16],
                    [16, 16],
                ],
            ],
            "expected": [
                684,
                9775,
                132388,
                23603659,
                276335856,
                482737486,
                928504131,
                221654878,
                900,
                204363,
                32550871,
                188481313,
                984133324,
                827733619,
                73337229,
                4754866,
                11380,
                24509011,
                262189558,
                77877831,
                85750856,
                775022885,
                16776,
                216907,
                28690693,
                342889162,
                848756209,
                908703893,
                42902734,
                265861897,
                306289573,
                449195744,
                22471,
                3323092,
                39095440,
                420276109,
                301135860,
                49,
                1078,
                13141,
                309545440,
                791932060,
                336366731,
                13284,
                177264,
                1994301,
                795814624,
                418,
                91228992,
                1204,
                198628,
                567,
                9540,
                152856,
                1974607,
                6642,
                92425,
                1257286,
                12942685,
                1035,
                15334,
                2233312,
                24475,
                254566,
                1264,
                13447,
                1,
            ],
        },
    ],
)
