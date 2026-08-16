import { runTests } from "../utils/utils.ts"

function dailyTemperatures(temperatures: number[]): number[] {
    const result: number[] = new Array(temperatures.length).fill(0)
    const stack: number[] = []

    stack.push(0)
    for (let i = 1; i < temperatures.length; i++) {
        let idx = stack.at(-1) as number
        while (temperatures[idx] < temperatures[i]) {
            result[idx] = i - idx
            stack.pop()
            idx = stack.at(-1) as number
        }
        stack.push(i)
    }

    return result
}

runTests(dailyTemperatures, [
    { input: [[73, 74, 75, 71, 69, 72, 76, 73]], expected: [1, 1, 4, 2, 1, 1, 0, 0] },
    { input: [[30, 40, 50, 60]], expected: [1, 1, 1, 0] },
    { input: [[30, 60, 90]], expected: [1, 1, 0] },
])
