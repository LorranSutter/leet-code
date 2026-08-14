import { runTests } from "../utils/utils.ts"

function missingInteger(nums: number[]): number {
    const nums_set = new Set(nums)
    const prefix_sum: number[] = new Array(nums.length + 1)

    prefix_sum[0] = 0
    prefix_sum[1] = nums[0]

    let start = 0
    let max_sum = prefix_sum[1]
    let min_missing_num = Infinity
    for (let i = 1; i < nums.length; i++) {
        prefix_sum[i + 1] = prefix_sum[i] + nums[i]

        if (nums[i] - 1 !== nums[i - 1]) {
            start = i
            if (max_sum < min_missing_num) {
                min_missing_num = max_sum
                while (nums_set.has(min_missing_num)) {
                    min_missing_num++
                }
            }
        }
        let diff_sum = prefix_sum[i + 1] - prefix_sum[start]
        if (diff_sum > max_sum) {
            max_sum = diff_sum
        }
    }

    if (max_sum < min_missing_num) {
        min_missing_num = max_sum
        while (nums_set.has(min_missing_num)) {
            min_missing_num++
        }
    }


    return min_missing_num
}

runTests(missingInteger, [
    { input: [[1, 2, 3, 2, 5]], expected: 6 },
    { input: [[3, 4, 5, 1, 12, 14, 13]], expected: 15 },
    { input: [[29, 30, 31, 32, 33, 34, 35, 36, 37]], expected: 297 },
    { input: [[37, 1, 2, 9, 5, 8, 5, 2, 9, 4]], expected: 38 },
    { input: [[14, 9, 6, 9, 7, 9, 10, 4, 9, 9, 4, 4]], expected: 15 },
])
