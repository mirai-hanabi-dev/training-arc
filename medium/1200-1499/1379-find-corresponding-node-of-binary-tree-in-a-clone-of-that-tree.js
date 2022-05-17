/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} original
 * @param {TreeNode} cloned
 * @param {TreeNode} target
 * @return {TreeNode}
 */

const getTargetCopy = (original, cloned, target) => {
  let result = null

  const dfs = node => {
    if (result || !node) {
      return
    }
    if (node.val == target.val) {
      result = node
      return
    }
    dfs(node.left)
    dfs(node.right)
  }
  dfs(cloned)

  return result
}
