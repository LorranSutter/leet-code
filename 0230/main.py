import sys
from typing import Optional
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests
from utils.utils import TreeNode, make_binary_tree_from_level_order


class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        smallest = 0
        count = 0

        def in_order(node: Optional[TreeNode]):
            nonlocal smallest, count

            if node == None:
                return

            in_order(node.left)

            count += 1
            if count == k:
                smallest = node.val
                return

            in_order(node.right)

        in_order(root)
        return smallest


root1 = make_binary_tree_from_level_order([3, 1, 4, None, 2])
root2 = make_binary_tree_from_level_order([5, 3, 6, 2, 4, None, None, 1])
run_tests(
    Solution().kthSmallest,
    [
        {"input": [root1, 1], "expected": 1},
        {"input": [root2, 3], "expected": 3},
    ],
)
