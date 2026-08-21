import { runTests } from "../utils/utils.ts"

/**
 * Preprocessing:
 * - Before starting any search, we run two cheap checks that answer the follow-up about search pruning:
 *   reject boards that can't possibly contain the word before paying for a single DFS call.
 * - If `word` is longer than the board has cells (`m * n`), it can never fit no matter what the letters
 *   are, so we return false immediately.
 * - Otherwise we count how many of each letter `word` needs, then walk the board once, decrementing that
 *   count every time a needed letter is seen and dropping the letter from the map once its count hits
 *   zero. If the map still has entries left after the walk, the board is short on at least one letter the
 *   word needs, and we return false without ever starting a search. This turns "the board can't have this
 *   word" cases into a single O(m * n + word.length) pass instead of discovering the same fact letter by
 *   letter, deep inside a doomed DFS branch.
 *
 * Search:
 * - The key insight is that this is backtracking on a grid: from every cell that matches the word's first
 *   letter, we walk the four neighbors trying to match the next letter, and undo (unmark) each step the
 *   moment a path dead-ends, so the same cells are free to be tried again from a different starting point
 *   or a different branch.
 * - `visited` tracks cells used by the *current* path only, encoded as `r * n + c` so we get a single set
 *   instead of a 2D array of booleans. A cell is added right before recursing into its neighbors and
 *   removed right after, which is what lets the same cell be reused across different attempted paths.
 * - `dfs` bails out immediately if the current cell's letter doesn't match `word[id_word]` or if the cell
 *   is already part of the path — that second check is what enforces "the same letter cell may not be
 *   used more than once" from the problem statement. If we're on the word's last letter and it matched,
 *   we're done: `found` is set and every pending call unwinds without doing more work.
 * - The outer double loop tries every cell as a possible starting point whose letter matches `word[0]`,
 *   and stops early (`if (found) return true`) as soon as one search succeeds instead of scanning the
 *   rest of the board. `visited.clear()` between attempts is only there for safety since a fully unwound
 *   `dfs` already empties it on its own.
 *
 *   worked example for board = [["A","B","C","E"],["S","F","E","S"],["A","D","E","E"]], word = "ABCESEEEFS":
 *
 *     A B C E
 *     S F E S
 *     A D E E
 *
 *   starting at (0,0)='A', the path that matches the whole word walks:
 *     (0,0)A (0,1)B (0,2)C (0,3)E (1,3)S (2,3)E (2,2)E (1,2)E (1,1)F (1,0)S
 *
 *   note the path revisits column 2 and 3's E's from different rows (1,2) and (2,2) — each is a distinct
 *   cell, so neither violates the "no reuse" rule even though the letter E repeats three times in the word.
 */
function exist(board: string[][], word: string): boolean {
    const m = board.length
    const n = board[0].length
    const word_length = word.length
    const visited = new Set<number>()
    let found = false

    // Word bigger than available letters in the board
    if (word.length > m * n) {
        return false
    }

    // Count how many of each letter the word needs
    const count_word = new Map<string, number>()
    for (const letter of word) {
        count_word.set(letter, (count_word.get(letter) ?? 0) + 1)
    }
    // Check if there are enough letters in the board
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            let letter = board[i][j]
            const currentValue = count_word.get(letter);
            if (currentValue !== undefined) {
                if (currentValue <= 1) {
                    count_word.delete(letter);
                } else {
                    count_word.set(letter, currentValue - 1);
                }
            }
        }
    }

    // There are not enough letters in the board for the word
    if (count_word.size > 0) {
        return false
    }

    function dfs(r: number, c: number, id_word: number) {
        if (board[r][c] !== word[id_word] || visited.has(r * n + c)) {
            return
        }
        if (id_word === word_length - 1) {
            found = true
            return
        }
        visited.add(r * n + c)

        // Up
        if (r > 0) {
            dfs(r - 1, c, id_word + 1)
        }
        // Down
        if (r < m - 1) {
            dfs(r + 1, c, id_word + 1)
        }
        // Right
        if (c < n - 1) {
            dfs(r, c + 1, id_word + 1)
        }
        // Left
        if (c > 0) {
            dfs(r, c - 1, id_word + 1)
        }

        visited.delete(r * n + c)
    }

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (board[i][j] === word[0]) {
                dfs(i, j, 0)
                if (found) {
                    return true
                }
                visited.clear()
            }
        }
    }

    return false
}

runTests(exist, [
    { input: [[["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]], "ABCCED"], expected: true },
    { input: [[["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]], "SEE"], expected: true },
    { input: [[["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]], "ABCB"], expected: false },
    { input: [[["a", "a"]], "aaa"], expected: false },
    { input: [[["a", "a", "a", "a"], ["a", "a", "a", "a"], ["a", "a", "a", "a"]], "aaaaaaaaaaaaa"], expected: false },
    { input: [[["a", "a"], ["a", "a"]], "aaaaa"], expected: false },
    { input: [[["A", "B", "C", "E"], ["S", "F", "E", "S"], ["A", "D", "E", "E"]], "ABCESEEEFS"], expected: true },
    { input: [[["A", "B", "C", "E"], ["S", "F", "E", "S"], ["A", "D", "E", "E"]], "ABCEFSADEESE"], expected: true },
])