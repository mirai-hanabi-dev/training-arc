class Solution:
    def longestValidParentheses(self, s):
        dp = [0] * len(s)

        def fetchOrZero(idx):
            if idx > 0:
                return dp[idx]
            return 0

        for idx in range(len(s)):
            if s[idx] == '(':
                continue
            if idx > 0 and s[idx - 1] == '(':
                dp[idx] = fetchOrZero(idx - 2) + 2
                continue
            if s[idx - 1] == ')':
                openIdx = idx - dp[idx - 1] - 1
                if openIdx >= 0 and s[openIdx] == '(':
                    dp[idx] = dp[idx - 1] + fetchOrZero(openIdx - 1) + 2

        result = 0
        for v in dp:
            result = max(result, v)

        return result


s = ''
print(Solution().longestValidParentheses(s))
