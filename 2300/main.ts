import { runTests } from "../utils/utils.ts"

function successfulPairs(spells: number[], potions: number[], success: number): number[] {
    potions.sort((a, b) => a - b)
    let result = new Array<number>(spells.length).fill(0)

    for (let i = 0; i < spells.length; i++) {
        let start = 0
        let end = potions.length - 1
        let mid = 0
        while (start <= end) {
            mid = start + Math.floor((end - start) / 2)
            if (spells[i] * potions[mid] < success) {
                start = mid + 1
            } else {
                end = mid - 1
            }
        }
        if (potions[mid] * spells[i] >= success) {
            result[i] = potions.length - mid
        } else if (potions[mid] * spells[i] < success) {
            result[i] = potions.length - mid - 1
        }
    }

    return result
}

runTests(successfulPairs, [
    { input: [[5, 1, 3], [1, 2, 3, 4, 5], 7], expected: [4, 0, 3] },
    { input: [[3, 1, 2], [8, 5, 8], 16], expected: [2, 0, 2] },
    { input: [[1, 2, 3, 4, 5, 6, 7], [1, 2, 3, 4, 5, 6, 7], 25], expected: [0, 0, 0, 1, 3, 3, 4] },
    {
        input: [
            [15, 39, 38, 35, 33, 25, 31, 12, 40, 27, 29, 16, 22, 24, 7, 36, 29, 34, 24, 9, 11, 35, 21, 3, 33, 10, 9, 27, 35, 17, 14, 3, 35, 35, 39, 23, 35, 14, 31, 7],
            [25, 19, 30, 37, 14, 30, 38, 22, 38, 38, 26, 33, 34, 23, 40, 28, 15, 29, 36, 39, 39, 37, 32, 38, 8, 17, 39, 20, 4, 39, 39, 7, 30, 35, 29, 23],
            317
        ],
        expected: [28, 33, 33, 33, 33, 33, 33, 23, 34, 33, 33, 29, 32, 33, 0, 33, 33, 33, 33, 13, 22, 33, 31, 0, 33, 17, 13, 33, 33, 30, 27, 0, 33, 33, 33, 33, 33, 27, 33, 0]
    },
])
