import { runTests } from "../utils/utils.ts"

function permute(nums: number[]): number[][] {
    const result: number[][] = [];
    const path: number[] = [];
    const used: boolean[] = new Array(nums.length).fill(false);
    const n = nums.length

    function backtrack() {
        // base case: path is a complete permutation
        if (path.length === n) {
            result.push([...path])
            return
        }

        // choice: loop over every index; skip if used
        for (let i = 0; i < n; i++) {
            //   choose: mark used, push to path
            if (used[i]) {
                continue
            }
            path.push(nums[i])
            used[i] = true
            //   explore: recurse
            backtrack()

            //   un-choose: pop, un-mark used
            path.pop()
            used[i] = false
        }
    }

    backtrack();
    return result;
}

runTests(permute, [
    { input: [[1, 2, 3]], expected: [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]] },
    { input: [[0, 1]], expected: [[0, 1], [1, 0]] },
    { input: [[1]], expected: [[1]] },
])
