import { runTests } from "../utils/utils.ts"

/**
 * - This was a nice one. The dummy solution would be to build an n x n matrix, iterate over each lamp, and
 *   flag every cell it illuminates. Then each query would just be an O(1) lookup into the matrix. Its only
 *   advantage was that O(1) lookup, though - n can be up to 1e9, so the matrix itself is way too much
 *   memory, and un-illuminating cells when a lamp gets turned off is a mess on top of that. This would get
 *   a TLE (or an OOM) for sure.
 * - The key insight is that we don't need to know which cells are illuminated - only whether a lamp exists
 *   anywhere in a given cell's row, column, or either diagonal. So instead of a full grid, we keep four
 *   hashmaps, one per direction, each mapping a row/column/diagonal index to how many lamps currently sit
 *   on it. A cell is illuminated whenever its row, column, main diagonal, or secondary diagonal has a count
 *   greater than zero in the matching map.
 * - The two diagonals need their own indexing scheme, since "row" and "column" aren't enough to identify
 *   them. Every cell on the same main diagonal (top-left to bottom-right) shares the same `row - col`
 *   value, and every cell on the same secondary diagonal (top-right to bottom-left) shares the same
 *   `row + col` value. That single number is what we use as the map key.
 *
 *   worked example: lamp at (0, 4) sets `sec_diag_on` count for key `0 + 4 = 4`. A query at (1, 3) checks
 *   `1 + 3 = 4` in that same map and finds it - the two cells sit on the same anti-diagonal, even though
 *   they share no row or column.
 *
 * - `lamp_set` dedupes the input lamp positions (the problem allows duplicates) and doubles as the "is this
 *   lamp currently on" check used later. `setIlluminated` increments or decrements all four maps for a
 *   given position, deleting a key once its count drops to zero so `.has()` checks stay accurate.
 * - Answering a query is just checking membership in the four maps for that cell's row, column, and two
 *   diagonal keys. Once a query is answered, `turnLampsOff` sweeps the 3x3 neighborhood around it (bounds
 *   checked against `n`) and, for every neighbor still in `lamp_set`, removes it from the set and decrements
 *   its contribution in `setIlluminated`.
 */


function gridIllumination(n: number, lamps: number[][], queries: number[][]): number[] {
    const lamp_set = new Set(lamps.map(l => `${l[0]},${l[1]}`))
    const query_result: number[] = new Array(queries.length).fill(0)
    const rows_on: Map<number, number> = new Map<number, number>()
    const cols_on: Map<number, number> = new Map<number, number>()
    const main_diag_on: Map<number, number> = new Map<number, number>()
    const sec_diag_on: Map<number, number> = new Map<number, number>()

    function setIlluminated(lamp: string, value: number) {
        const [r, c] = lamp.split(",").map(Number)
        let count = rows_on.get(r) || 0
        rows_on.set(r, count + value);
        if (rows_on.get(r) == 0) {
            rows_on.delete(r)
        }

        count = cols_on.get(c) || 0
        cols_on.set(c, count + value);
        if (cols_on.get(c) == 0) {
            cols_on.delete(c)
        }

        count = main_diag_on.get(r - c) || 0
        main_diag_on.set(r - c, count + value)
        if (main_diag_on.get(r - c) == 0) {
            main_diag_on.delete(r - c)
        }

        count = sec_diag_on.get(r + c) || 0
        sec_diag_on.set(r + c, count + value)
        if (sec_diag_on.get(r + c) == 0) {
            sec_diag_on.delete(r + c)
        }
    }

    function turnLampsOff(pos: number[], n: number, lamp_set: Set<string>) {
        for (let dr = -1; dr <= 1; dr++) {
            for (let dc = -1; dc <= 1; dc++) {
                const r = pos[0] + dr
                const c = pos[1] + dc
                if (r < 0 || r >= n || c < 0 || c >= n) {
                    continue
                }
                const idx = `${r},${c}`
                if (lamp_set.has(idx)) {
                    lamp_set.delete(idx)
                    setIlluminated(idx, -1)
                }
            }
        }
    }

    for (const lamp of lamp_set) {
        setIlluminated(lamp, 1)
    }

    for (let i = 0; i < queries.length; i++) {
        if (rows_on.has(queries[i][0])) {
            query_result[i] = 1
            turnLampsOff(queries[i], n, lamp_set)
        } else if (cols_on.has(queries[i][1])) {
            query_result[i] = 1
            turnLampsOff(queries[i], n, lamp_set)
        } else if (main_diag_on.has(queries[i][0] - queries[i][1])) {
            query_result[i] = 1
            turnLampsOff(queries[i], n, lamp_set)
        } else if (sec_diag_on.has(queries[i][0] + queries[i][1])) {
            query_result[i] = 1
            turnLampsOff(queries[i], n, lamp_set)
        }
    }

    return query_result
}

runTests(gridIllumination, [
    { input: [5, [[0, 0], [4, 4]], [[1, 1], [1, 0]]], expected: [1, 0] },
    { input: [5, [[0, 0], [4, 4]], [[1, 1], [1, 1]]], expected: [1, 1] },
    { input: [5, [[0, 0], [0, 4]], [[0, 4], [0, 1], [1, 4]]], expected: [1, 1, 0] },
    { input: [5, [[0, 0], [4, 4], [3, 1]], [[0, 4], [0, 1], [1, 4], [1, 3]]], expected: [1, 1, 1, 1] },
    { input: [6, [[2, 5], [4, 2], [0, 3], [0, 5], [1, 4], [4, 2], [3, 3], [1, 0]], [[4, 3], [3, 1], [5, 3], [0, 5], [4, 4], [3, 3]]], expected: [1, 0, 1, 1, 0, 1] },
])
