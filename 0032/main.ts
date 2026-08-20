import { runTests } from "../utils/utils.ts"

function longestValidParentheses(s: string): number {
    let longest = 0

    let left = 0
    let right = 0
    let count = 0
    for (const p of s) {
        if (p === "(") {
            left++
        } else {
            right++
        }

        if (left === right) {
            longest = Math.max(longest, 2 * right)
        }
        if (right > left) {
            left = 0
            right = 0
        }
    }

    left = 0
    right = 0
    count = 0
    for (let i = s.length - 1; i >= 0; i--) {
        if (s[i] === "(") {
            left++
        } else {
            right++
        }

        if (left === right) {
            longest = Math.max(longest, 2 * left)
        }
        if (left > right) {
            left = 0
            right = 0
        }
    }

    return longest
};

runTests(longestValidParentheses, [
    { input: ["(()"], expected: 2 },
    { input: [")()())"], expected: 4 },
    { input: [""], expected: 0 },
    { input: ["()"], expected: 2 },
    { input: ["())()()((()()()))()()()()()()())))(()(())))()())()()()())(((()))))()()())()()()()())))("], expected: 28 },
])
