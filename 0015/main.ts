import { runTests } from "../utils/utils.ts"

function threeSum(nums: number[]): number[][] {
    nums.sort((a, b) => a - b)
    let result = new Set<string>()

    let left, right, s
    for (let i = 0; i < nums.length - 2; i++) {
        left = i + 1
        right = nums.length - 1
        while (left < right) {
            s = nums[i] + nums[left] + nums[right]
            if (s === 0) {
                result.add(JSON.stringify([nums[i], nums[left], nums[right]]))
            }
            if (s > 0) {
                right--
            } else {
                left++
            }
        }
    }

    return Array.from(result).map((str) => JSON.parse(str))
}

runTests(threeSum, [
    { input: [[-1, 0, 1, 2, -1, -4]], expected: [[-1, -1, 2], [-1, 0, 1]] },
    { input: [[0, 1, 1]], expected: [] },
    { input: [[0, 0, 0]], expected: [[0, 0, 0]] },
    { input: [[-1, 0, 1, 0]], expected: [[-1, 0, 1]] },
    { input: [[-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6]], expected: [[-4, -2, 6], [-4, 0, 4], [-4, 1, 3], [-4, 2, 2], [-2, -2, 4], [-2, 0, 2]] },
])
