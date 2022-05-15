/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
const dfs = (node, currentDepth, result) => {
  if (!node) {
    return
  }
  if (currentDepth > result.depth) {
    result.depth = currentDepth
    result.sum = node.val
  } else if (currentDepth == result.depth) {
    result.sum += node.val
  }

  dfs(node.left, currentDepth + 1, result)
  dfs(node.right, currentDepth + 1, result)
}

var deepestLeavesSum = function(root) {
  let result = {
    depth: 0,
    sum: 0
  }
  dfs(root, 1, result)
  return result.sum
};
