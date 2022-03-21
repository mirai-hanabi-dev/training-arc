class Solution:
    def isInteger(self, s):
        hasNumber = False
        for idx, char in enumerate(s):
            if (char == '+' or char == '-'):
                if idx == 0:
                    continue
                return False
            if char < '0' or char > '9':
                return False
            else:
                hasNumber = True
        return hasNumber

    def isDecimal(self, s):
        dotPos = None
        startNumPos = None
        for idx, char in enumerate(s):
            if char == '+' or char == '-':
                if idx != 0:
                    return False
                else:
                    continue
            if char == '.':
                if dotPos is not None:
                    return False
                if startNumPos is None:
                    startNumPos = idx
                dotPos = idx
                continue
            if char < '0' or char > '9':
                return False
            elif startNumPos is None:
                startNumPos = idx

        if dotPos is None:
            return False

        return dotPos - startNumPos != 0 or len(s) - 1 - dotPos != 0

    def isNumber(self, s):
        ePos = None
        for idx, char in enumerate(s):
            if char.lower() == 'e':
                if ePos is not None:
                    return False
                ePos = idx

        if ePos is None:
            return self.isDecimal(s) or self.isInteger(s)
        return (self.isDecimal(s[:ePos]) or self.isInteger(s[:ePos])) \
            and self.isInteger(s[ePos + 1:])
