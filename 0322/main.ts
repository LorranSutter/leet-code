import { runTests } from "../utils/utils.ts"

function coinChange(coins: number[], amount: number): number {
    const dp: number[] = new Array(amount + 1).fill(Infinity);
    dp[0] = 0;

    for (let a = 1; a <= amount; a++) {
        for (const coin of coins) {
            if (coin <= a) {
                dp[a] = Math.min(dp[a], dp[a - coin] + 1)
            }
        }
    }

    return dp[amount] === Infinity ? -1 : dp[amount];
}

runTests(coinChange, [
    { input: [[1, 2, 5, 10], 18], expected: 4 },
    { input: [[186, 419, 83, 408], 6249], expected: 20 },
    { input: [[1, 2, 5], 11], expected: 3 },
    { input: [[2], 3], expected: -1 },
    { input: [[1], 0], expected: 0 },
])
