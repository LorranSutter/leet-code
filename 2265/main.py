import sys
from typing import Tuple
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests, TreeNode, make_binary_tree_from_level_order


class Solution:
    def averageOfSubtree(self, root: TreeNode) -> int:
        count = 0

        def dfs(node: TreeNode) -> Tuple[int, int]:
            nonlocal count
            if node == None:
                return 0, 0

            left_sum, left_count = dfs(node.left)
            right_sum, right_count = dfs(node.right)

            numerator = left_sum + node.val + right_sum
            num_nodes = left_count + 1 + right_count

            avg = numerator // num_nodes
            if avg == node.val:
                count += 1

            return numerator, num_nodes

        dfs(root)
        return count


root1 = make_binary_tree_from_level_order([4, 8, 5, 0, 1, None, 6])
root2 = make_binary_tree_from_level_order([1])
root3 = make_binary_tree_from_level_order([1, None, 3, None, 1, None, 3])
run_tests(
    Solution().averageOfSubtree,
    [
        {"input": [root1], "expected": 5},
        {"input": [root2], "expected": 1},
        {"input": [root3], "expected": 1},
    ],
)
