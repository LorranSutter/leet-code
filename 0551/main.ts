import { runTests } from "../utils/utils.ts"

function checkRecord(s: string): boolean {
    let absent = 0
    let late = 0

    for (const att of s) {
        switch (att) {
            case 'A':
                absent++
                late = 0
                break;
            case 'L':
                late++
                break;
            default:
                late = 0
                break;
        }

        if (absent >= 2 || late >= 3) {
            return false
        }
    }

    return true
}

runTests(checkRecord, [
    { input: ["PPALLP"], expected: true },
    { input: ["PPALLL"], expected: false },
    { input: ["PAALPL"], expected: false },
    { input: ["ALLAPPL"], expected: false },
])
