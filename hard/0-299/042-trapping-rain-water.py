class Solution:
    def trap(self, height):
        n = len(height)
        maxLeft = [0 for _ in range(n)]
        maxRight = [0 for _ in range(n)]

        for i in range(n):
            maxLeft[i] = height[i] if i == 0 \
                else max(height[i], maxLeft[i - 1])
            maxRight[n - i - 1] = height[n - i - 1] if i == 0 \
                else max(height[n - i - 1], maxRight[n - i])

        res = 0
        for i in range(1, n - 1):
            res += min(maxLeft[i], maxRight[i]) - height[i]

        return res
