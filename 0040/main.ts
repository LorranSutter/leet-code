import { runTests } from "../utils/utils.ts"

function combinationSum2(candidates: number[], target: number): number[][] {
    const result: number[][] = []
    candidates.sort((a, b) => a - b)

    let comb: number[] = []
    function backtrack(start: number, sum: number) {
        if (sum == target) {
            result.push([...comb])
            return
        }
        if (sum > target || start >= candidates.length) {
            return
        }

        for (let i = start; i < candidates.length; i++) {
            if (i > start && candidates[i] === candidates[i - 1]) continue

            comb.push(candidates[i])
            backtrack(i + 1, sum + candidates[i])
            comb.pop()
        }
    }

    backtrack(0, 0)
    return result
}

runTests(combinationSum2, [
    { input: [[10, 1, 2, 7, 6, 1, 5], 8], expected: [[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]] },
    { input: [[2, 5, 2, 1, 2], 5], expected: [[1, 2, 2], [5]] },
])