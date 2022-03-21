class Solution:
    def getPermutation(self, n, k):
        factorials = []
        for i in range(n):
            if i == 0:
                factorials.append(1)
            else:
                factorials.append(i * factorials[i - 1])

        result = []
        used = [False for _ in range(n)]

        def dfs(n, k):
            if n == 0:
                return
            nextNum = k // factorials[n - 1] + \
                (1 if k % factorials[n - 1] != 0 else 0)

            count = 0
            for i in range(len(used)):
                if used[i]:
                    continue
                count += 1
                if nextNum == 0 or count == nextNum:
                    used[i] = True
                    result.append(str(i + 1))
                    dfs(n - 1, k - (nextNum - 1) * factorials[n - 1])
                    return

        dfs(n, k)

        return ''.join(result)
