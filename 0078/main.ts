import { runTests } from "../utils/utils.ts"

/**
 * Solution 1:
 * - The key insight is that every subset of `nums` maps to a unique bitmask from 0 to 2^n - 1, where bit
 *   `j` says whether `nums[j]` is included. Counting from 0 to 2^n - 1 in binary is the same as walking
 *   through every possible in/out combination for the n elements, so we never have to think recursively
 *   about it at all.
 * - We preallocate `result` with `2^n` slots (all sharing the placeholder `[]` from `.fill([])`), then
 *   loop `i` from 1 to `2^n - 1` and, for each `j`, check whether bit `j` of `i` is set to decide whether
 *   `nums[j]` belongs in that subset. Index 0 is left untouched, which is fine since it's supposed to be
 *   the empty subset anyway.
 *
 *   worked example for nums = [1, 2, 3], i = 5:
 *     5 in binary is 101
 *     bit 0 set -> include nums[0] = 1
 *     bit 1 clear -> skip nums[1] = 2
 *     bit 2 set -> include nums[2] = 3
 *     subset = [1, 3]
 *
 * Solution 2:
 * - Same problem, solved by walking a decision tree instead of counting bitmasks: at each index we
 *   either take `nums[i]` or leave it out, then recurse on `i + 1`. Once `i` runs past the end of `nums`,
 *   every element has been decided one way or the other, so the current `subset` is a complete subset and
 *   gets copied into `result`.
 * - The two recursive calls per index are what give us both branches: `backtrack(i + 1)` right after
 *   pushing `nums[i]` covers "included", and `backtrack(i + 1)` again after popping it covers "excluded".
 *
 *   Obs: because this walks the decision tree in a different order than the bitmask counts up, the
 *   subsets come out in a different order than solution 1's (and than the problem's sample output), so the
 *   solution 2 tests use `unorderedSubsetsEqual` to compare results regardless of order.
 */

// Bit manipulation solution
function subsets1(nums: number[]): number[][] {
    const size = 2 ** nums.length
    const result = new Array(2 ** nums.length).fill([])

    for (let i = 1; i < size; i++) {
        let new_subset = []
        for (let j = 0; j < nums.length; j++) {
            if ((i & (1 << j)) != 0) {
                new_subset.push(nums[j])
            }
        }
        result[i] = new_subset
    }

    return result
}

// Backtracking solution
function subsets2(nums: number[]): number[][] {
    const result: number[][] = []

    const subset: number[] = []
    function backtrack(i: number) {
        if (i >= nums.length) {
            result.push([...subset])
            return
        }

        subset.push(nums[i])
        backtrack(i + 1)

        subset.pop()
        backtrack(i + 1)

    }

    backtrack(0)
    return result
}

runTests(subsets1, [
    { input: [[1, 2, 3]], expected: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]] },
    { input: [[0]], expected: [[], [0]] },
])

function unorderedSubsetsEqual(actual: number[][], expected: number[][]): boolean {
    if (actual.length !== expected.length) {
        return false
    }

    const normalize = (subsets: number[][]) =>
        subsets.map(subset => [...subset].sort((a, b) => a - b).join(",")).sort()

    const a = normalize(actual)
    const b = normalize(expected)

    return a.every((val, i) => val === b[i])
}

runTests(subsets2, [
    { input: [[1, 2, 3]], expected: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]] },
    { input: [[0]], expected: [[], [0]] },
], unorderedSubsetsEqual)
