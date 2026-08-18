import { runTests } from "../utils/utils.ts"

function combinationSum(candidates: number[], target: number): number[][] {
    const result: number[][] = []
    candidates.sort((a, b) => a - b)

    let sum = 0
    let subset: number[] = []
    function backtrack(i: number) {
        if (sum == target) {
            result.push([...subset])
            return
        }
        if (sum > target || i >= candidates.length) {
            return
        }

        sum += candidates[i]
        subset.push(candidates[i])
        backtrack(i)

        let last = subset.pop()
        if (last) sum -= last
        backtrack(i + 1)
    }

    backtrack(0)
    return result
}

runTests(combinationSum, [
    { input: [[2, 3, 6, 7], 7], expected: [[2, 2, 3], [7]] },
    { input: [[6, 7, 2, 3], 7], expected: [[2, 2, 3], [7]] },
    { input: [[2, 3, 5], 8], expected: [[2, 2, 2, 2], [2, 3, 3], [3, 5]] },
    { input: [[2], 1], expected: [] },
])
