import sys
from typing import List, Optional
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests, TreeNode, make_binary_tree_from_level_order


class Solution:
    def getAllElements(
        self, root1: Optional[TreeNode], root2: Optional[TreeNode]
    ) -> List[int]:
        def in_order(node: TreeNode, arr: List[int]):
            if node == None:
                return

            in_order(node.left, arr)
            arr.append(node.val)
            in_order(node.right, arr)

        arr1, arr2 = [], []
        in_order(root1, arr1)
        in_order(root2, arr2)

        return sorted(arr1 + arr2)


root1_1 = make_binary_tree_from_level_order([2, 1, 4])
root1_2 = make_binary_tree_from_level_order([1, 0, 3])
root2_1 = make_binary_tree_from_level_order([1, None, 8])
root2_2 = make_binary_tree_from_level_order([8, 1])
run_tests(
    Solution().getAllElements,
    [
        {"input": [root1_1, root1_2], "expected": [0, 1, 1, 2, 3, 4]},
        {"input": [root2_1, root2_2], "expected": [1, 1, 8, 8]},
    ],
)
