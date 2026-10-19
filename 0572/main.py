import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests, TreeNode, make_binary_tree_from_level_order


class Solution:
    def isSubtree(self, root: TreeNode | None, subRoot: TreeNode | None) -> bool:
        def is_subtree(root1: TreeNode, root2: TreeNode) -> bool:
            if root1 == root2 == None:
                return True
            if root1 == None and root2 != None:
                return False
            if root1 != None and root2 == None:
                return False
            if root1.val != root2.val:
                return False
            return is_subtree(root1.left, root2.left) and is_subtree(
                root1.right, root2.right
            )

        def dfs(root1: TreeNode, root2: TreeNode) -> bool:
            if root2 == None:
                return True
            if root1 == None and root2 != None:
                return False
            if root1.val == root2.val:
                if is_subtree(root1, root2):
                    return True
            return dfs(root1.left, root2) or dfs(root1.right, root2)

        return dfs(root, subRoot)


root1 = make_binary_tree_from_level_order([3, 4, 5, 1, 2])
subRoot1 = make_binary_tree_from_level_order([4, 1, 2])
root2 = make_binary_tree_from_level_order([3, 4, 5, 1, 2, None, None, None, None, 0])
subRoot2 = make_binary_tree_from_level_order([4, 1, 2])
run_tests(
    Solution().isSubtree,
    [
        {"input": [root1, subRoot1], "expected": True},
        {"input": [root2, subRoot2], "expected": False},
    ],
)
