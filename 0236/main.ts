import { runTests, TreeNode, makeBinaryTreeFromLevelOrder } from "../utils/utils.ts"


function lowestCommonAncestor(root: TreeNode | null, p: TreeNode, q: TreeNode): TreeNode | null {
    if (!root || root.val === p.val || root.val === q.val) {
        return root
    }

    let left = lowestCommonAncestor(root.left, p, q)
    let right = lowestCommonAncestor(root.right, p, q)

    if (left && right) {
        return root
    }

    return left || right
}

const root1 = makeBinaryTreeFromLevelOrder([3, 5, 1, 6, 2, 0, 8, null, null, 7, 4], null) as TreeNode
const root2 = makeBinaryTreeFromLevelOrder([3, 5, 1, 6, 2, 0, 8, null, null, 7, 4], null) as TreeNode
const root3 = makeBinaryTreeFromLevelOrder([1, 2], null) as TreeNode
runTests(lowestCommonAncestor, [
    { input: [root1, new TreeNode(5), new TreeNode(1)], expected: new TreeNode(3) },
    { input: [root2, new TreeNode(5), new TreeNode(4)], expected: new TreeNode(5) },
    { input: [root3, new TreeNode(1), new TreeNode(2)], expected: new TreeNode(1) },
], (actual, expected) => actual?.val === expected?.val)
