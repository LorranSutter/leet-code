import sys
from typing import Optional
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests, TreeNode, make_binary_tree_from_level_order


class Solution:
    def checkTree(self, root: Optional[TreeNode]) -> bool:
        return root.val == root.left.val + root.right.val


root1 = make_binary_tree_from_level_order([10, 4, 6])
root2 = make_binary_tree_from_level_order([5, 3, 1])
run_tests(
    Solution().checkTree,
    [
        {"input": [root1], "expected": True},
        {"input": [root2], "expected": False},
    ],
)
