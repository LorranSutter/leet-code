/**
 * - The key insight is that this is a coin-change problem in disguise: the "coins" are perfect squares
 *   (1, 4, 9, 16, ...) up to n, and we want the fewest of them that add up to n.
 * - We build a `dp` array where `dp[i]` is the least number of perfect squares summing to `i`, with
 *   `dp[0] = 0` as the base case (zero squares needed to make zero).
 * - For every `i` from 1 to n, we try every square that fits (square <= i) and take whichever one leaves
 *   the cheapest remainder: `dp[i] = min(dp[i], dp[i - square] + 1)`. The `+ 1` accounts for the square
 *   we just used.
 * - Because squares are generated in increasing order, we can break out of the inner loop as soon as a
 *   square exceeds i, instead of checking the rest.
 *
 *   worked example for n = 12, squares = [0, 1, 4, 9]:
 *     dp[1] = dp[0] + 1 = 1                  (1)
 *     dp[2] = dp[1] + 1 = 2                  (1+1)
 *     dp[3] = dp[2] + 1 = 3                  (1+1+1)
 *     dp[4] = min(dp[3]+1, dp[0]+1) = 1      (4)
 *     ...
 *     dp[8] = min(dp[7]+1, dp[4]+1) = 2      (4+4)
 *     dp[12] = min(dp[11]+1, dp[8]+1, dp[3]+1) = 3   (4+4+4)
 *

 */

function numSquares(n: number): number {
    const dp: number[] = new Array(n + 1).fill(Infinity);
    const squares: number[] = new Array(Math.floor(Math.sqrt(n)))
    for (let i = 0; i <= Math.floor(Math.sqrt(n)) + 1; i++) {
        squares[i] = i * i
    }

    dp[0] = 0

    for (let i = 1; i <= n; i++) {
        for (const square of squares) {
            if (square > i) {
                break
            }
            dp[i] = Math.min(dp[i], dp[i - square] + 1)
        }
    }

    return dp[n]
}

console.log(numSquares(12) === 3);
console.log(numSquares(13) === 2);
