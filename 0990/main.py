import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests

"""
Solution:

- The idea here lies on noticing something interesting: we only have two kinds of equations, == and !=,
  and equality is transitive. If a == b and b == c, then by the transitive property a == c, so any number
  we pick for a works for b and c too, as long as we pick the same one for all three.
- Now the twist: if we add an inequality like a != b, it breaks the group above outright. But a != c
  breaks it just as badly, even though c never shows up next to a directly - because a == c already
  follows from the equalities we already have.
- In summary, a, b and c belong to the same group, so any inequality that disturbs the group makes the
  whole set of equations false.
- This logic can be translated into a Union-Find problem, where the parents are all the letters of the
  alphabet, and the groups are the sets of variables tied together by ==.
- We evaluate each equation in order. If it's an equality (==), we call `union` to merge its two letters
  into the same group. If it's an inequality (!=), we can't check it yet - the groups aren't finished
  forming - so we just stash it in an array to evaluate later.
- Once every equality has built the groups, we walk back through the stashed inequalities and check
  whether each one disturbs a group, i.e. whether its two letters ended up in the same set. That's just
  `find(a) == find(b)`: if it's true for any inequality, the equations contradict each other and we
  return False; if none do, every inequality sits between separate groups and the assignment is possible.
"""


class Solution:
    def equationsPossible(self, equations: List[str]) -> bool:
        parent = list(range(26))

        def find(v: int):
            if parent[v] != v:
                parent[v] = find(parent[v])
            return parent[v]

        def union(v1: int, v2: int):
            parent[find(v1)] = find(v2)

        # ord('a') -> 97
        inequalities = []
        for eq in equations:
            if "==" in eq:
                v1, v2 = eq.split("==")
                union(ord(v1) - 97, ord(v2) - 97)
            else:
                v1, v2 = eq.split("!=")
                inequalities.append((ord(v1) - 97, ord(v2) - 97))

        for v1, v2 in inequalities:
            if find(v1) == find(v2):
                return False

        return True


run_tests(
    Solution().equationsPossible,
    [
        {"input": [["a==b", "b!=a"]], "expected": False},
        {"input": [["a==b", "b!=b"]], "expected": False},
        {"input": [["b==a", "a==b"]], "expected": True},
        {"input": [["a==b", "b!=c", "c==a"]], "expected": False},
        {"input": [["a==b", "e==c", "b==c", "a!=e"]], "expected": False},
        {"input": [["c==c", "f!=a", "f==b", "b==c"]], "expected": True},
        {"input": [["e==d", "e==a", "f!=d", "b!=c", "a==b"]], "expected": True},
        {"input": [["e==d", "e==a", "f!=d", "e!=b", "a==b"]], "expected": False},
    ],
)
