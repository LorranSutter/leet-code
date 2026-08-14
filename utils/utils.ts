import { isDeepStrictEqual } from "node:util"

type TestCase<Args extends unknown[], R> = {
    input: Args
    expected: R
}

export function runTests<Args extends unknown[], R>(
    fn: (...args: Args) => R,
    cases: TestCase<Args, R>[]
): void {
    const failures: { index: number; input: Args; expected: R; actual: R }[] = []

    cases.forEach((testCase, index) => {
        const actual = fn(...testCase.input)
        if (!isDeepStrictEqual(actual, testCase.expected)) {
            failures.push({ index, input: testCase.input, expected: testCase.expected, actual })
        }
    })

    if (failures.length === 0) {
        console.log(`✅ All ${cases.length} test case(s) passed`)
        return
    }

    console.log(`❌ ${failures.length}/${cases.length} test case(s) failed`)
    for (const failure of failures) {
        console.log(`  Case ${failure.index}:`)
        console.log(`    Input:    ${JSON.stringify(failure.input)}`)
        console.log(`    Expected: ${JSON.stringify(failure.expected)}`)
        console.log(`    Actual:   ${JSON.stringify(failure.actual)}`)
    }
}
