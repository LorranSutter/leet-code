import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests, TreeNode, make_binary_tree_from_level_order


class Solution:
    def isSameTree(self, p: TreeNode | None, q: TreeNode | None) -> bool:
        def dfs(node1: TreeNode, node2: TreeNode) -> bool:
            if node1 == None and node2 != None:
                return False
            if node1 != None and node2 == None:
                return False
            if node1 == None and node2 == None:
                return True
            if node1.val != node2.val:
                return False

            return dfs(node1.left, node2.left) and dfs(node1.right, node2.right)

        return dfs(p, q)


root1_1 = make_binary_tree_from_level_order([1, 2, 3])
root2_1 = make_binary_tree_from_level_order([1, 2, 3])
root1_2 = make_binary_tree_from_level_order([1, 2])
root2_2 = make_binary_tree_from_level_order([1, None, 2])
root1_3 = make_binary_tree_from_level_order([1, 2, 1])
root2_3 = make_binary_tree_from_level_order([1, 1, 2])
run_tests(
    Solution().isSameTree,
    [
        {"input": [root1_1, root2_1], "expected": True},
        {"input": [root1_2, root2_2], "expected": False},
        {"input": [root1_3, root2_3], "expected": False},
    ],
)
