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
 * @return {number[][]}
 */
const levelOrder = (root) => {
  const result = [];
  let queue = [];
  if (root) {
    queue.push(root);
  }
  while (queue.length > 0) {
    result.push(queue.map((node) => node.val));
    const nextQueue = [];
    for (const node of queue) {
      if (node.left) {
        nextQueue.push(node.left);
      }
      if (node.right) {
        nextQueue.push(node.right);
      }
    }
    queue = nextQueue;
  }
  return result;
};
