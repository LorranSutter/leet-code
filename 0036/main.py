import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        m = len(board)

        # Check rows and columns
        for i in range(m):
            seen_row = set()
            seen_col = set()
            for j in range(m):
                if board[i][j] != "." and board[i][j] in seen_row:
                    return False
                if board[j][i] != "." and board[j][i] in seen_col:
                    return False
                seen_row.add(board[i][j])
                seen_col.add(board[j][i])

        # Check sub-boxes
        for offset_i in range(3):
            for offset_j in range(3):
                seen = set()
                for i in range(3 * offset_i, 3 + 3 * offset_i):
                    for j in range(3 * offset_j, 3 + 3 * offset_j):
                        if board[i][j] == ".":
                            continue
                        if board[i][j] in seen:
                            return False
                        seen.add(board[i][j])

        return True


run_tests(
    Solution().isValidSudoku,
    [
        {
            "input": [
                [
                    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
                    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
                    [".", "9", "8", ".", ".", ".", ".", "6", "."],
                    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
                    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
                    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
                    [".", "6", ".", ".", ".", ".", "2", "8", "."],
                    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
                    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
                ]
            ],
            "expected": True,
        },
        {
            "input": [
                [
                    ["8", "3", ".", ".", "7", ".", ".", ".", "."],
                    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
                    [".", "9", "8", ".", ".", ".", ".", "6", "."],
                    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
                    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
                    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
                    [".", "6", ".", ".", ".", ".", "2", "8", "."],
                    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
                    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
                ]
            ],
            "expected": False,
        },
        {
            "input": [
                [
                    [".", ".", "4", ".", ".", ".", "6", "3", "."],
                    [".", ".", ".", ".", ".", ".", ".", ".", "."],
                    ["5", ".", ".", ".", ".", ".", ".", "9", "."],
                    [".", ".", ".", "5", "6", ".", ".", ".", "."],
                    ["4", ".", "3", ".", ".", ".", ".", ".", "1"],
                    [".", ".", ".", "7", ".", ".", ".", ".", "."],
                    [".", ".", ".", "5", ".", ".", ".", ".", "."],
                    [".", ".", ".", ".", ".", ".", ".", ".", "."],
                    [".", ".", ".", ".", ".", ".", ".", ".", "."],
                ]
            ],
            "expected": False,
        },
    ],
)
