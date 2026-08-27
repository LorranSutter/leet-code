import { runTests } from "../utils/utils.ts"

function decodeString(s: string): string {
    const numStack: number[] = []
    const strStack: string[] = []
    let current = ""
    let num = 0

    for (const ch of s) {
        if (ch >= "0" && ch <= "9") {
            num = num * 10 + Number(ch)
        } else if (ch === "[") {
            numStack.push(num)
            strStack.push(current)
            num = 0
            current = ""
        } else if (ch === "]") {
            const repeat = numStack.pop()!
            current = strStack.pop()! + current.repeat(repeat)
        } else {
            current += ch
        }
    }

    return current
}

runTests(decodeString, [
    { input: ["3[a]2[bc]"], expected: "aaabcbc" },
    { input: ["3[a2[c]]"], expected: "accaccacc" },
    { input: ["2[abc]3[cd]ef"], expected: "abcabccdcdcdef" },
    { input: ["ef3[a2[c]]xy"], expected: "efaccaccaccxy" },
    { input: ["c3[2[a]1[b]]d"], expected: "caabaabaabd" },
    { input: ["11[2[a]]"], expected: "aaaaaaaaaaaaaaaaaaaaaa" },
    { input: ["2[3[a2[c]]d]e"], expected: "accaccaccdaccaccaccde" },
    { input: ["3[z]2[2[y]pq4[2[jk]e1[f]]]ef"], expected: "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef" },
])

