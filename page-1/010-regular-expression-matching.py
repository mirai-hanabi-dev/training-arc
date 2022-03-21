class Solution:
    def isMatch(self, s, p):
        memo = {}

        # Check if s[i:] is match with p[j:]
        def dp(i, j):
            if (i, j) in memo:
                return memo[(i, j)]

            subS = s[i:]
            subP = p[j:]
            if len(subP) > 1 and subP[1] == '*':
                if subP[0] == '.':
                    isMatch = False
                    for _ in range(i, len(s) + 1):
                        if dp(_, j + 2):
                            isMatch = True
                            break
                    memo[(i, j)] = isMatch
                else:
                    isMatch = dp(i, j + 2)
                    char = subP[0]
                    for _ in range(i, len(s)):
                        if s[_] == char and dp(_ + 1, j + 2):
                            isMatch = True
                            break
                        elif s[_] != char:
                            break
                    memo[(i, j)] = isMatch
            elif len(subP) == 0:
                memo[(i, j)] = True if len(subS) == 0 else False
            elif len(subS) == 0:
                memo[(i, j)] = False
            elif subP[0] == '.' or subS[0] == subP[0]:
                memo[(i, j)] = dp(i + 1, j + 1)
            else:
                memo[(i, j)] = False

            return memo[(i, j)]

        dp(0, 0)

        return memo[(0, 0)]
