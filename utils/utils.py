from typing import Any, Callable


def run_tests(fn: Callable[..., Any], cases: list[dict[str, Any]]) -> None:
    failures = []

    for index, case in enumerate(cases):
        actual = fn(*case["input"])
        if actual != case["expected"]:
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
