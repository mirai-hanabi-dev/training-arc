# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def __init__(self):
        self.minTreeCache = dict()
        self.maxTreeCache = dict()

    def isValidBST(self, root):
        if root is None:
            return True

        if root.left and self.maxTree(root.left) >= root.val:
            return False

        if root.right and self.minTree(root.right) <= root.val:
            return False

        return self.isValidBST(root.left) and self.isValidBST(root.right)

    def minTree(self, node):
        if node is None:
            return float('Inf')

        if id(node) in self.minTreeCache:
            return self.minTreeCache[id(node)]

        res = min(self.minTree(node.left),
                  self.minTree(node.right),
                  node.val)
        self.minTreeCache[id(node)] = res
        return res

    def maxTree(self, node):
        if node is None:
            return -float('Inf')

        if id(node) in self.maxTreeCache:
            return self.maxTreeCache[id(node)]

        res = max(self.maxTree(node.left),
                  self.maxTree(node.right),
                  node.val)
        self.maxTreeCache[id(node)] = res
        return res


node1 = TreeNode(1)
node5 = TreeNode(5)
node4 = TreeNode(4)
node3 = TreeNode(3)
node6 = TreeNode(6)

node5.left = node1
node5.right = node4
node4.left = node3
node4.right = node6

print(Solution().isValidBST(node5))
