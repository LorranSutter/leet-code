import { runTests } from "../utils/utils.ts"

function letterCombinations(digits: string): string[] {
    const result: string[] = []
    const n = digits.length
    const phoneRecord: Record<string, string> = {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz"
    }

    let comb = ""
    function backtrack(i: number) {
        if (comb.length === n) {
            result.push(comb)
            return
        }

        for (const letter of phoneRecord[digits[i]]) {
            comb += letter
            backtrack(i + 1)
            comb = comb.slice(0, comb.length - 1)
        }
    }

    backtrack(0)
    return result
}

runTests(letterCombinations, [
    { input: ["23"], expected: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"] },
    { input: ["2"], expected: ["a", "b", "c"] },
    { input: ["9"], expected: ["w", "x", "y", "z"] },
])
