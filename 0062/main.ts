import { runTests } from "../utils/utils.ts"

/*
- The key insight is that the only way to arrive at any cell (i, j) is from directly above it or
  directly to its left, since the robot only ever moves down or right. That means the number of unique
  paths to (i, j) is just the sum of the paths to those two neighbors: dp[i][j] = dp[i-1][j] + dp[i][j-1].

- The base case is the first row and the first column. A cell on row 1 can only be reached by moving
  right the whole way, and a cell on column 1 can only be reached by moving down the whole way, so
  there's exactly one path to every cell on either edge: dp[i][1] = dp[1][j] = 1.

- From there we fill the table in increasing order of i and j, since dp[i][j] only depends on values
  above and to the left of it, which are already computed by the time we get there.

  Here's the table for m = 3, n = 3 (rows are i, columns are j):
      j=1 j=2 j=3
  i=1   1   1   1
  i=2   1   2   3
  i=3   1   3   6
  dp[3][3] = dp[2][3] + dp[3][2] = 3 + 3 = 6, matching the expected output for that test case.
*/

function uniquePaths(m: number, n: number): number {
    const dp = new Map<string, number>()

    for (let i = 1; i <= m; i++) {
        dp.set(`${i},${1}`, 1)
    }
    for (let j = 1; j <= n; j++) {
        dp.set(`${1},${j}`, 1)
    }

    for (let i = 2; i <= m; i++) {
        for (let j = 2; j <= n; j++) {
            dp.set(`${i},${j}`, (dp.get(`${i - 1},${j}`) ?? 0) + (dp.get(`${i},${j - 1}`) ?? 0))
        }
    }

    return dp.get(`${m},${n}`) ?? 0
}

runTests(uniquePaths, [
    { input: [3, 7], expected: 28 },
    { input: [3, 2], expected: 3 },
    { input: [3, 3], expected: 6 },
    { input: [18, 6], expected: 26334 },
    { input: [21, 21], expected: 137846528820 },
])
