# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def inOrderTraversal(self, node):
        if not node:
            return
        if len(self.inOrder) == self.k:
            return
        self.inOrderTraversal(node.left)
        if len(self.inOrder) < self.k:
            self.inOrder.append(node.val)
        self.inOrderTraversal(node.right)

    def kthSmallest(self, root, k):
        self.inOrder = []
        self.k = k
        self.inOrderTraversal(root)
        return self.inOrder[-1]


node1 = TreeNode(1)
node2 = TreeNode(2)
node3 = TreeNode(3)
node4 = TreeNode(4)

node3.left = node1
node3.right = node4
node1.right = node2

print(Solution().kthSmallest(node3, 1))
