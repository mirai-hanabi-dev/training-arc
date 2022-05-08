class Solution:
    def __init__(self):
        self.cache = {}

    def isScramble(self, u, v):
        if (u, v) in self.cache:
            return self.cache[(u, v)]

        answer = None
        if u == v:
            answer = True
        elif sorted(u) != sorted(v):
            answer = False
        else:
            answer = False
            for i in range(1, len(u)):
                if self.isScramble(u[:i], v[:i]) \
                        and self.isScramble(u[i:], v[i:]):
                    answer = True
                    break

                if self.isScramble(u[:i], v[-i:]) \
                        and self.isScramble(u[i:], v[:-i]):
                    answer = True
                    break

        self.cache[(u, v)] = answer
        return answer
