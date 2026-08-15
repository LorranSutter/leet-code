import { runTests } from "../utils/utils.ts"

function minFlips(a: number, b: number, c: number): number {
    if ((a | b) === c) {
        return 0
    }

    const size = Math.max(a, b, c).toString(2).length

    let count = 0
    let a_bit, b_bit, c_bit
    for (let i = size; i >= 0; i--) {
        a_bit = a >> i & 1
        b_bit = b >> i & 1
        c_bit = c >> i & 1

        if ((a_bit | b_bit) == c_bit) {
            continue
        }

        count++
        if (c_bit == 0 && a_bit == 1 && b_bit == 1) {
            count++
        }
    }

    return count
}

runTests(minFlips, [
    { input: [2, 6, 5], expected: 3 },
    { input: [4, 2, 7], expected: 1 },
    { input: [1, 2, 3], expected: 0 },
])
