from typing import Any, Callable


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
