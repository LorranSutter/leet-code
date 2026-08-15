import { runTests } from "../utils/utils.ts"

function isPerfectSquare1(num: number): boolean {
    for (let i = 1; ; i++) {
        let square = i * i
        if (square > num) {
            return false
        }
        if (square == num) {
            return true
        }
    }
};

function isPerfectSquare2(num: number): boolean {
    let left = 1
    let middle = 0
    let right = num
    let square = 0

    while (left <= right) {
        middle = Math.floor(left + (right - left) / 2)
        square = middle * middle
        if (square == num) {
            return true
        }
        if (square < num) {
            left = middle + 1
        } else {
            right = middle - 1
        }
    }

    return false
};

runTests(isPerfectSquare1, [
    { input: [16], expected: true },
    { input: [14], expected: false },
    { input: [1], expected: true },
])

runTests(isPerfectSquare2, [
    { input: [16], expected: true },
    { input: [14], expected: false },
    { input: [1], expected: true },
])
