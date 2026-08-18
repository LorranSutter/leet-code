import { runTests } from "../utils/utils.ts"

function permuteUnique(nums: number[]): number[][] {
    nums.sort((a, b) => a - b)
    const result: number[][] = []
    const used: boolean[] = new Array(nums.length).fill(false)
    const path: number[] = []
    const n = nums.length

    function backtrack() {
        if (path.length === n) {
            result.push([...path])
            return
        }

        for (let i = 0; i < n; i++) {
            if (used[i]) {
                continue
            }
            if (i > 0 && nums[i] === nums[i - 1] && !used[i - 1]) {
                continue
            }

            path.push(nums[i])
            used[i] = true
            backtrack()

            path.pop()
            used[i] = false
        }

    }

    backtrack()
    return result
}

runTests(permuteUnique, [
    { input: [[1, 1, 2]], expected: [[1, 1, 2], [1, 2, 1], [2, 1, 1]] },
    { input: [[1, 2, 3]], expected: [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]] },
    { input: [[3, 3, 0, 3]], expected: [[0, 3, 3, 3], [3, 0, 3, 3], [3, 3, 0, 3], [3, 3, 3, 0]] },
])
