import { runTests } from "../utils/utils.ts"

// Sieve of Eratosthenes
function countPrimes(n: number): number {
    const primes: boolean[] = new Array(n).fill(true)
    primes[0] = false
    primes[1] = false

    let p = 2
    while (p * p <= n) {
        if (primes[p]) {
            for (let i = p * p; i < n; i += p) {
                primes[i] = false
            }
        }
        p++
    }

    return primes.reduce((acc, curr) => acc + (curr ? 1 : 0), 0)
}

runTests(countPrimes, [
    { input: [10], expected: 4 },
    { input: [0], expected: 0 },
    { input: [1], expected: 0 },
    { input: [2], expected: 0 },
])
