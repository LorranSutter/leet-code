import sys
from typing import List, Optional
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests
from utils.utils import TreeNode, make_binary_tree_from_level_order


class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        max_diameter = 0

        def diameter(node: TreeNode) -> int:
            nonlocal max_diameter
            if node == None:
                return 0

            left = diameter(node.left)
            right = diameter(node.right)

            max_diameter = max(max_diameter, left + right)

            return 1 + max(left, right)

        diameter(root)
        return max_diameter


root1 = make_binary_tree_from_level_order([1, 2, 3, 4, 5])
root2 = make_binary_tree_from_level_order([1, 2])
root3 = make_binary_tree_from_level_order([1])
run_tests(
    Solution().diameterOfBinaryTree,
    [
        {"input": [root1], "expected": 3},
        {"input": [root2], "expected": 1},
        {"input": [root3], "expected": 0},
    ],
)
