import { runTests } from "../utils/utils.ts"

function threeSumClosest(nums: number[], target: number): number {
    nums.sort((a, b) => a - b)

    let left, right, s, diff
    let closest_sum = 0
    let closest_value = Infinity
    for (let i = 0; i < nums.length - 2; i++) {
        left = i + 1
        right = nums.length - 1
        while (left < right) {
            s = nums[i] + nums[left] + nums[right]
            diff = Math.abs(s - target)
            if (diff < closest_value) {
                closest_value = diff
                closest_sum = s
            }
            if (s > target) {
                right--
            } else {
                left++
            }
        }
    }

    return closest_sum
}

runTests(threeSumClosest, [
    { input: [[-1, 2, 1, -4], 1], expected: 2 },
    { input: [[0, 0, 0], 1], expected: 0 },
    { input: [[1, 1, 1, 0], -100], expected: 2 },
    { input: [[4, 0, 5, -5, 3, 3, 0, -4, -5], -2], expected: -2 },
    { input: [[10, 20, 30, 40, 50, 60, 70, 80, 90], 1], expected: 60 },
])
