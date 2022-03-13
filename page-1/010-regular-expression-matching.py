# Only faster than 5% T.T

class Solution:
    def isMatch(self, s, p):
        if len(s) == 0 and len(p) == 0:
            return True
        if len(s) > 0 and len(p) == 0:
            return False

        if len(p) >= 2 and p[1] == '*':
            for i in range(len(s) + 1):
                if len(set(s[:i])) > 1 and p[0] != '.':
                    continue
                if len(set(s[:i])) == 1 and p[0] != '.' and p[0] != s[0]:
                    continue
                if self.isMatch(s[i:], p[2:]):
                    return True
            return False

        if len(s) == 0:
            return False

        if p[0] == '.' or s[0] == p[0]:
            return self.isMatch(s[1:], p[1:])

        return False


s = "ab"
p = ".*c"
assert Solution().isMatch(s, p) is False
