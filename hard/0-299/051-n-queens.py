# Old submission

class Solution:
    def isValidPosition(self, posY, currentState, n):
        if posY in currentState:
            return False
        posX = len(currentState)
        for x in range(len(currentState)):
            y = currentState[x]
            if posX - x == posY - y:
                return False
            if posX - y == x - posY:
                return False
        return True

    def choose(self, currentState):
        if len(currentState) == self.n:
            self.solutions.append(currentState)
            return
        for i in range(self.n):
            if self.isValidPosition(i, currentState, self.n):
                nextState = currentState + [i]
                self.choose(nextState)

    def solveNQueens(self, n):
        self.n = n
        self.solutions = []

        self.choose([])

        stringSolutions = []
        for solution in self.solutions:
            stringSolution = []
            for pos in solution:
                stringSolution.append('.' * pos + 'Q' + '.' * (n - pos - 1))
            stringSolutions.append(stringSolution)

        return stringSolutions
