# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def sortedArrayToBST(self, nums):
        if len(nums) == 0:
            return None

        left = 0
        right = len(nums) - 1
        mid = left + (right - left + 1) // 2

        node = TreeNode(nums[mid])
        node.left = self.sortedArrayToBST(nums[left:mid])
        node.right = self.sortedArrayToBST(nums[mid+1:right+1])

        return node


nums = [-10, -3, 0, 1, 5, 9]
node = Solution().sortedArrayToBST(nums)
