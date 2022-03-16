import collections


class Solution:
    def findSubstring(self, s, words):
        wordLength = len(words[0])
        wordsCount = collections.Counter(words)

        def sliding_window(left):
            wordsFound = collections.Counter()
            wordsUsed = 0
            right = left
            while True:
                right += wordLength
                if right > len(s):
                    break

                currentWord = s[right - wordLength:right]
                if currentWord not in wordsCount:
                    # Current word not in bag of words.
                    # Move to the next window.
                    wordsFound = collections.Counter()
                    wordsUsed = 0
                    left = right
                    continue

                # Current word in bag of words.
                wordsFound[currentWord] += 1
                wordsUsed += 1
                while wordsFound[currentWord] > wordsCount[currentWord]:
                    firstWord = s[left:left + wordLength]
                    wordsFound[firstWord] -= 1
                    wordsUsed -= 1
                    left += wordLength

                if wordsUsed == len(words):
                    answer.append(left)

        answer = []
        for i in range(wordLength):
            sliding_window(i)
        return answer


s = "barfoothefoobarman"
words = ["bar", "foo"]
print(Solution().findSubstring(s, words))
