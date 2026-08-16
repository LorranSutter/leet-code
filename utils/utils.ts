import { isDeepStrictEqual } from "node:util"

export class TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.left = (left === undefined ? null : left)
        this.right = (right === undefined ? null : right)
    }
}

export function makeBinaryTree(nums: (number | null)[]): TreeNode | null {
    if (!nums || nums.length == 0) {
        return null
    }

    const root = new TreeNode(nums[0] as number)

    for (let i = 1; i < nums.length; i++) {
        addBinaryTreeNode(root, nums[i])
    }

    return root
}

export function printTree(root: TreeNode | null) {
    if (!root) {
        return
    }

    process.stdout.write(`${root.val} `);
    printTree(root.left)
    printTree(root.right)
}

// Add TreeNode to the binary tree
function addBinaryTreeNode(root: TreeNode, val: number | null) {
    if (!val) {
        return
    }
    if (val < root.val) {
        if (root.left === null) {
            root.left = new TreeNode(val)
        } else {
            addBinaryTreeNode(root.left, val)
        }
    } else {
        if (root.right === null) {
            root.right = new TreeNode(val)
        } else {
            addBinaryTreeNode(root.right, val)
        }
    }
}

type TestCase<Args extends unknown[], R> = {
    input: Args
    expected: R
}

export function runTests<Args extends unknown[], R>(
    fn: (...args: Args) => R,
    cases: TestCase<Args, R>[]
): void {
    const failures: { index: number; input: Args; expected: R; actual: R }[] = []

    cases.forEach((testCase, index) => {
        const actual = fn(...testCase.input)
        if (!isDeepStrictEqual(actual, testCase.expected)) {
            failures.push({ index, input: testCase.input, expected: testCase.expected, actual })
        }
    })

    if (failures.length === 0) {
        console.log(`✅ All ${cases.length} test case(s) passed`)
        return
    }

    console.log(`❌ ${failures.length}/${cases.length} test case(s) failed`)
    for (const failure of failures) {
        console.log(`  Case ${failure.index}:`)
        console.log(`    Input:    ${JSON.stringify(failure.input)}`)
        console.log(`    Expected: ${JSON.stringify(failure.expected)}`)
        console.log(`    Actual:   ${JSON.stringify(failure.actual)}`)
    }
}
