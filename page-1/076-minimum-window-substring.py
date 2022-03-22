import collections


class Solution:
    def minWindow(self, s, t):
        counterS = collections.Counter()
        counterT = collections.Counter(t)

        optimalLength = len(s) + 1
        optimalPos = (None, None)

        # Instead of repeatly calculate if the current string is valid,
        # we should use a counter to count number of satisfy characters.
        # => validSubstring <=> counter == len(counterT.keys())
        def validSubstring():
            for key in counterT.keys():
                if counterS[key] < counterT[key]:
                    return False
            return True

        left = 0
        right = -1
        while right + 1 < len(s):
            right += 1
            counterS[s[right]] += 1
            while validSubstring() and left <= right:
                if right - left + 1 < optimalLength:
                    optimalLength = right - left + 1
                    optimalPos = (left, right)
                counterS[s[left]] -= 1
                left += 1

        if optimalLength == len(s) + 1:
            return ''

        left, right = optimalPos
        return s[left:right + 1]
