import { runTests } from "../utils/utils.ts"

function eraseOverlapIntervals(intervals: number[][]): number {
    intervals.sort((a, b) => a[1] - b[1])

    let count = 1
    let end = intervals[0][1]
    for (let i = 0; i < intervals.length; i++) {
        if (intervals[i][0] >= end) {
            end = intervals[i][1]
            count++
        }
    }

    return intervals.length - count
}

runTests(eraseOverlapIntervals, [
    { input: [[[1, 2], [2, 3], [3, 4], [1, 3]]], expected: 1 },
    { input: [[[1, 2], [1, 2], [1, 2]]], expected: 2 },
    { input: [[[1, 2], [2, 3]]], expected: 0 },
    { input: [[[1, 100], [11, 22], [1, 11], [2, 12]]], expected: 2 },
])
