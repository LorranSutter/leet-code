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

console.log(countPrimes(10) === 4);
console.log(countPrimes(0) === 0);
console.log(countPrimes(1) === 0);
console.log(countPrimes(2) === 0);
