from typing import Any, Callable


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def make_binary_tree(nums: list[int]) -> TreeNode | None:
    if not nums:
        return None

    root = TreeNode(nums[0])

    for val in nums[1:]:
        _add_binary_tree_node(root, val)

    return root


def _add_binary_tree_node(root: TreeNode, val: int) -> None:
    if val < root.val:
        if root.left is None:
            root.left = TreeNode(val)
        else:
            _add_binary_tree_node(root.left, val)
    else:
        if root.right is None:
            root.right = TreeNode(val)
        else:
            _add_binary_tree_node(root.right, val)


def make_binary_tree_from_level_order(
    nums: list[int], null_value: int = None
) -> TreeNode | None:
    if not nums or nums[0] == null_value:
        return None

    root = TreeNode(nums[0])
    queue = [root]
    i = 1

    while queue and i < len(nums):
        current = queue.pop(0)

        if i < len(nums) and nums[i] != null_value:
            current.left = TreeNode(nums[i])
            queue.append(current.left)
        i += 1

        if i < len(nums) and nums[i] != null_value:
            current.right = TreeNode(nums[i])
            queue.append(current.right)
        i += 1

    return root


def _build_tree_string(root: TreeNode | None) -> tuple[list[str], int, int, int]:
    if root is None:
        return [], 0, 0, 0

    line1: list[str] = []
    line2: list[str] = []
    node_repr = str(root.val)

    new_root_width = gap_size = len(node_repr)

    l_box, l_box_width, l_root_start, l_root_end = _build_tree_string(root.left)
    r_box, r_box_width, r_root_start, r_root_end = _build_tree_string(root.right)

    if l_box_width > 0:
        l_root = (l_root_start + l_root_end) // 2 + 1
        line1.append(" " * (l_root + 1))
        line1.append("_" * (l_box_width - l_root))
        line2.append(" " * l_root + "/")
        line2.append(" " * (l_box_width - l_root))
        new_root_start = l_box_width + 1
        gap_size += 1
    else:
        new_root_start = 0

    line1.append(node_repr)
    line2.append(" " * new_root_width)

    if r_box_width > 0:
        r_root = (r_root_start + r_root_end) // 2
        line1.append("_" * r_root)
        line1.append(" " * (r_box_width - r_root + 1))
        line2.append(" " * r_root + "\\")
        line2.append(" " * (r_box_width - r_root - 1))
        gap_size += 1

    new_root_end = new_root_start + new_root_width - 1

    gap = " " * gap_size
    new_box = ["".join(line1), "".join(line2)]
    for i in range(max(len(l_box), len(r_box))):
        l_line = l_box[i] if i < len(l_box) else " " * l_box_width
        r_line = r_box[i] if i < len(r_box) else " " * r_box_width
        new_box.append(l_line + gap + r_line)

    return new_box, len(new_box[0]), new_root_start, new_root_end


def print_tree(root: TreeNode | None) -> None:
    lines, *_ = _build_tree_string(root)
    print("\n".join(lines))


def print_matrix(matrix: list[list[Any]], separator: str = " ") -> None:
    for row in matrix:
        print(separator.join(str(value) for value in row))


def equal_unordered(s1: list[Any], s2: list[Any]) -> bool:
    if len(s1) != len(s2):
        return False

    remaining = list(s2)
    for value in s1:
        try:
            remaining.remove(value)
        except ValueError:
            return False

    return True


def equal_matrix_unordered(m1: list[list[Any]], m2: list[list[Any]]) -> bool:
    def canonical(matrix: list[list[Any]]) -> list[tuple[Any, ...]]:
        return sorted(tuple(sorted(row)) for row in matrix)

    return canonical(m1) == canonical(m2)


def run_tests(
    fn: Callable[..., Any],
    cases: list[dict[str, Any]],
    is_equal: Callable[[Any, Any], bool] = lambda actual, expected: actual == expected,
) -> None:
    failures = []

    for index, case in enumerate(cases):
        actual = fn(*case["input"])
        if not is_equal(actual, case["expected"]):
            failures.append((index, case["input"], case["expected"], actual))

    if not failures:
        print(f"✅ All {len(cases)} test case(s) passed")
        return

    print(f"❌ {len(failures)}/{len(cases)} test case(s) failed")
    for index, input_, expected, actual in failures:
        print(f"  Case {index}:")
        print(f"    Input:    {input_}")
        print(f"    Expected: {expected}")
        print(f"    Actual:   {actual}")
