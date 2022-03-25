class Solution:
    def largestRectangleArea(self, heights):
        # left[i] holds the idx of the left most bar
        # which bars from left[i] to i have greater or equal
        # height with bar ith.
        n = len(heights)

        left = [i for i in range(n)]
        stack = []
        for i in range(n):
            while len(stack) and heights[stack[len(stack) - 1]] >= heights[i]:
                stack.pop()
            if len(stack) == 0:
                left[i] = 0
            else:
                left[i] = stack[len(stack) - 1] + 1
            stack.append(i)

        stack = []
        right = [i for i in range(n)]
        for i in range(n - 1, -1, -1):
            while len(stack) and heights[stack[len(stack) - 1]] >= heights[i]:
                stack.pop()
            if len(stack) == 0:
                right[i] = n - 1
            else:
                right[i] = stack[len(stack) - 1] - 1
            stack.append(i)

        res = 0
        for i in range(n):
            res = max(res, (right[i] - left[i] + 1) * heights[i])

        return res
