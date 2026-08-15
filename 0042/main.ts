import { runTests } from "../utils/utils.ts"

function trap(height: number[]): number {
    let count = 0
    let left = 0
    let right = height.length - 1
    let max_left = height[left]
    let max_right = height[right]

    while (left < right) {
        if (max_left > max_right) {
            right--
            if (height[right] < max_right) {
                count += max_right - height[right]
            } else {
                max_right = Math.max(max_right, height[right])
            }
        } else {
            left++
            if (height[left] < max_left) {
                count += max_left - height[left]
            } else {
                max_left = Math.max(max_left, height[left])
            }
        }
    }
    return count
}

runTests(trap, [
    { input: [[0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]], expected: 6 },
    { input: [[4, 2, 0, 3, 2, 5]], expected: 9 },
    { input: [[3, 4, 1, 2, 2, 5, 1, 0, 2]], expected: 10 },
])
