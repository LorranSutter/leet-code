import { runTests } from "../utils/utils.ts"

function uniformArray(nums1: number[]): boolean {
    return true
};

runTests(uniformArray, [
    { input: [[2, 3]], expected: true },
    { input: [[4, 6]], expected: true },
    { input: [[1, 3]], expected: true },
    { input: [[1, 2]], expected: true },
])
