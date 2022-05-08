BOARD_LENGTH = 9
BOX_LENGTH = 3
BOX_TOP_LEFT = [
    [0, 0],
    [0, 3],
    [0, 6],
    [3, 0],
    [3, 3],
    [3, 6],
    [6, 0],
    [6, 3],
    [6, 6],
]


class Solution:
    def solveSudoku(self, board):
        # Strategy:
        # 1. Loop through the number 1->9.
        # 2. Find cells which able to fill with that number.
        # 3. Repeat till no empty cells left.

        countEmptyCells = 0

        for i in range(BOARD_LENGTH):
            for j in range(BOARD_LENGTH):
                if board[i][j] == '.':
                    countEmptyCells += 1

        while True:
            if countEmptyCells == 0:
                break

            isFilled = False
            for i in range(BOARD_LENGTH):
                countFillCells = self.greedy(board, str(i + 1))

                countEmptyCells -= countFillCells
                if countEmptyCells == 0:
                    break
                if countFillCells > 0:
                    isFilled = True

            # The above strategy sometimes doesn't find possible solution.
            # Break to fallback to backtracking.
            if not isFilled:
                break

        self.dfs(board)

    def dfs(self, board):
        def isSafeRow(strNum, i):
            for _ in range(BOARD_LENGTH):
                if board[i][_] == strNum:
                    return False
            return True

        def isSafeColumn(strNum, j):
            for _ in range(BOARD_LENGTH):
                if board[_][j] == strNum:
                    return False
            return True

        def isSafeBox(strNum, i, j):
            boxIndex = (i // 3) * 3 + (j // 3)
            boxi, boxj = BOX_TOP_LEFT[boxIndex]
            for i in range(BOX_LENGTH):
                for j in range(BOX_LENGTH):
                    if board[boxi + i][boxj + j] == strNum:
                        return False
            return True

        for i in range(BOARD_LENGTH):
            for j in range(BOARD_LENGTH):
                if board[i][j] == '.':
                    for num in range(BOARD_LENGTH):
                        strNum = str(num + 1)
                        if isSafeRow(strNum, i) and isSafeColumn(strNum, j) \
                                and isSafeBox(strNum, i, j):
                            board[i][j] = strNum
                            if self.dfs(board):
                                return True
                            else:
                                board[i][j] = '.'
                    return False
        return True

    def greedy(self, board, num):
        check = [[False for _ in range(BOARD_LENGTH)]
                 for _ in range(BOARD_LENGTH)]

        def fillCheck(i, j):
            for _ in range(BOARD_LENGTH):
                check[i][_] = True
                check[_][j] = True

        for i in range(BOARD_LENGTH):
            for j in range(BOARD_LENGTH):
                if board[i][j] != '.':
                    check[i][j] = True
                if board[i][j] == num:
                    fillCheck(i, j)

        countFillCells = 0
        for boxi, boxj in BOX_TOP_LEFT:
            posi, posj = None, None
            shouldWriteNum = True
            for i in range(BOX_LENGTH):
                if not shouldWriteNum:
                    break
                for j in range(BOX_LENGTH):
                    if board[boxi + i][boxj + j] == num:
                        shouldWriteNum = False
                        break
                    if check[boxi + i][boxj + j] is False:
                        if posi is not None and posj is not None:
                            shouldWriteNum = False
                            break
                        posi = boxi + i
                        posj = boxj + j

            if shouldWriteNum and posi is not None and posj is not None:
                board[posi][posj] = num
                countFillCells += 1

        return countFillCells
