# Old solution
class Solution:
    def calculate(self, s):
        s = s.replace(' ', '')

        def calculate_without_parentheses(s):
            s = s.replace('--', '+')
            res = 0
            remind = 0
            s = '+' + s + '+'
            operation = '+'
            for char in s:
                if char == '+' or char == '-':
                    if operation == '+':
                        res += remind
                    else:
                        res -= remind
                    operation = char
                    remind = 0
                else:
                    remind = 10 * remind + int(char)
            return res

        while True:
            close_paren_idx = s.find(')')
            if close_paren_idx == -1:
                break
            open_paren_idx = s[:close_paren_idx].rfind('(')
            next_s = s[open_paren_idx + 1:close_paren_idx]
            res = calculate_without_parentheses(next_s)
            s = s[:open_paren_idx] + str(res) + s[close_paren_idx + 1:]

        return calculate_without_parentheses(s)
