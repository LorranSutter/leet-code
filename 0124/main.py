import sys
import math
from typing import Optional
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests
from utils.utils import TreeNode, make_binary_tree_from_level_order

"""
Solution:

- The key insight is that a path can bend at most once, at a single "turning point" node, because once
  it bends both of that node's sides are used up and no further node can be attached without visiting
  something twice. So for every node we only need two numbers: the best path that bends here (using
  both children), and the best straight branch this node can hand up to its parent (using at most one
  child, since the parent can only continue the path in one direction).

- We do a post-order DFS. For each node we first ask its left and right children how much they can
  contribute if the path walked into them and kept going straight, clamping negative contributions to
  0 with `max(0, ...)` - a path is never made worse by skipping a branch that would only subtract from
  the sum.

- With `left` and `right` in hand, `root.val + left + right` is the best path that bends at this node,
  and we compare it against the running `max_sum`. That variable starts at `-inf` rather than 0, so a
  tree that's entirely negative (e.g. a single node `[-3]`) still reports its own value instead of a
  fake 0.

- What we `return` to the caller is different from what we compare against `max_sum`: we can only hand
  the parent a straight branch, `root.val + max(left, right)`, since attaching both children here would
  make it impossible for the grandparent to extend the path without reusing this node.

  Here's how it plays out on example 2, `[-10, 9, 20, null, null, 15, 7]`:

           -10
           /  \\
          9    20
              /  \\
            15    7

  - `dfs(9)`  -> leaf, `left = right = 0`, bend value `9`, `max_sum` becomes `9`, returns `9`
  - `dfs(15)` -> leaf, bend value `15`, `max_sum` becomes `15`, returns `15`
  - `dfs(7)`  -> leaf, bend value `7`, `max_sum` stays `15`, returns `7`
  - `dfs(20)` -> `left = max(0, 15) = 15`, `right = max(0, 7) = 7`, bend value `20 + 15 + 7 = 42`,
    `max_sum` becomes `42`, returns `20 + max(15, 7) = 35`
  - `dfs(-10)` -> `left = max(0, 9) = 9`, `right = max(0, 35) = 35`, bend value `-10 + 9 + 35 = 34`,
    which doesn't beat `42`, so `max_sum` stays `42`; it returns `-10 + max(9, 35) = 25`, which is never
    used since `-10` is the root

  The final answer is `42`, matching the `15 -> 20 -> 7` path - note that this path never touches the
  root at all, which is exactly why `max_sum` has to be tracked globally rather than returned from the
  top-level call.
"""


class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        max_sum = -math.inf

        def dfs(root: TreeNode) -> int:
            nonlocal max_sum
            if root == None:
                return 0

            # Ignore negative increments
            left = max(0, dfs(root.left))
            right = max(0, dfs(root.right))

            max_sum = max(max_sum, root.val + left + right)

            return root.val + max(left, right)

        dfs(root)
        return max_sum


root1 = make_binary_tree_from_level_order([1, 2, 3])
root2 = make_binary_tree_from_level_order([-10, 9, 20, None, None, 15, 7])
root3 = make_binary_tree_from_level_order([1, 1, 1, 1, 1, 1, 1])
root4 = make_binary_tree_from_level_order([2, -1, -2])
root5 = make_binary_tree_from_level_order([1, -2, 3])
root6 = make_binary_tree_from_level_order([-3])
root7 = make_binary_tree_from_level_order(
    [5, 4, 8, 11, None, 13, 4, 7, 2, None, None, None, 1]
)
root8 = make_binary_tree_from_level_order(
    [9, 6, -3, None, None, -6, 2, None, None, 2, None, -6, -6, -6]
)

run_tests(
    Solution().maxPathSum,
    [
        {"input": [root1], "expected": 6},
        {"input": [root2], "expected": 42},
        {"input": [root3], "expected": 5},
        {"input": [root4], "expected": 2},
        {"input": [root5], "expected": 4},
        {"input": [root6], "expected": -3},
        {"input": [root7], "expected": 48},
        {"input": [root8], "expected": 16},
    ],
)
