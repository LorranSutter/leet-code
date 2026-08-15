import { runTests } from "../utils/utils.ts"

function rob(nums: number[]): number {
    let prev = 0;
    let curr = 0;

    for (const num of nums) {
        const next = Math.max(curr, prev + num);
        prev = curr;
        curr = next;
    }

    return curr;
}

runTests(rob, [
    { input: [[1, 2, 3, 1]], expected: 4 },
    { input: [[2, 7, 9, 3, 1]], expected: 12 },
])
