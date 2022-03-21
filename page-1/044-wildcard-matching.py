class Solution:
    def isMatch(self, s, p):
        newP = ''
        for idx in range(len(p)):
            if idx < len(p) - 1 and p[idx] == '*' and p[idx + 1] == '*':
                continue
            else:
                newP += p[idx]
        p = newP

        memo = [[None for _ in range(len(p) + 1)] for _ in range(len(s) + 1)]

        # Check if s[i:] is match with p[j:]
        def dp(i, j):
            if memo[i][j] is not None:
                return memo[i][j]

            if j < len(p) and p[j] == '*':
                memo[i][j] = False
                for _ in range(i, len(s) + 1):
                    if dp(_, j + 1):
                        memo[i][j] = True
                        break
            elif j == len(p):
                memo[i][j] = True if i == len(s) else False
            elif i == len(s):
                memo[i][j] = False
            elif p[j] == '?' or s[i] == p[j]:
                memo[i][j] = dp(i + 1, j + 1)
            else:
                memo[i][j] = False

            return memo[i][j]

        dp(0, 0)

        return memo[0][0]
