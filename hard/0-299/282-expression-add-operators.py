# Old solution
class Solution:
    res = []
    num = None
    target = None

    def recursive(self, num, equation, value, lastNumber):
        if not num:
            if value == self.target:
                self.res.append(equation)
            return

        # Heuristic
        if len(num) == 1 or num[0] != '0':
            possibleExtremumValues = [
                value + int(num),
                value - int(num),
                value - lastNumber + lastNumber * int(num)
            ]
            if self.target < min(possibleExtremumValues) \
                    or self.target > max(possibleExtremumValues):
                return

        for i in range(1, len(num) + 1):
            nextNumber = num[:i]
            if len(nextNumber) > 1 and nextNumber[0] == '0':
                break

            self.recursive(
                num[i:],
                equation + '+' + nextNumber,
                value + int(nextNumber),
                int(nextNumber))

            self.recursive(
                num[i:],
                equation + '-' + nextNumber,
                value - int(nextNumber),
                -int(nextNumber))

            self.recursive(
                num[i:],
                equation + '*' + nextNumber,
                value - lastNumber + lastNumber * int(nextNumber),
                lastNumber * int(nextNumber))

    def addOperators(self, num, target):
        self.res = []
        self.num = num
        self.target = target

        for i in range(1, len(num) + 1):
            nextNumber = num[:i]
            if len(nextNumber) > 1 and nextNumber[0] == '0':
                break
            self.recursive(
                num[i:],
                nextNumber,
                int(nextNumber),
                int(nextNumber))

        return self.res
