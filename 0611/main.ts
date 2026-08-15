import { runTests } from "../utils/utils.ts"

function triangleNumber(nums: number[]): number {
    nums.sort((a, b) => a - b)

    let count = 0
    for (let i = nums.length - 1; i >= 2; i--) {
        let left = 0
        let right = i - 1
        while (left < right) {
            if (nums[left] + nums[right] > nums[i]) {
                count += right - left
                right--
            } else {
                left++
            }
        }
    }

    return count
}

runTests(triangleNumber, [
    { input: [[2, 2, 3, 4]], expected: 3 },
    { input: [[4, 2, 3, 4]], expected: 4 },
    { input: [[11, 4, 9, 6, 15, 18]], expected: 10 },
])
