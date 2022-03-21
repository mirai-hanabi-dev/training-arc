class Solution:
    def lineJustify(self, words, maxWidth, isLastLine=False):
        currentLength = 0
        for word in words:
            currentLength += len(word)

        spaceLeft = maxWidth - currentLength
        if isLastLine:
            res = ' '.join(words)
            return res + ' ' * (maxWidth - len(res))

        if len(words) == 1:
            words.append('')

        numberOfGroups = len(words) - 1

        # Now we need to divide spaceLeft spaces
        # into numberOfGroups groups.
        div = spaceLeft // numberOfGroups
        mod = spaceLeft % numberOfGroups

        res = ''
        for idx, word in enumerate(words):
            res += word
            if idx < len(words) - 1:
                res += ' ' * div + (' ' if mod > 0 else '')
                mod -= 1

        return res

    def fullJustify(self, words, maxWidth):
        result = []
        currentLength = 0
        numberOfWords = 0
        for idx, word in enumerate(words):
            if currentLength + len(word) + numberOfWords <= maxWidth:
                currentLength += len(word)
                numberOfWords += 1
            else:
                result.append(self.lineJustify(
                    words[idx - numberOfWords:idx], maxWidth))
                currentLength = len(word)
                numberOfWords = 1
            if idx == len(words) - 1:
                result.append(self.lineJustify(
                    words[idx - numberOfWords + 1:],
                    maxWidth,
                    isLastLine=True))
        return result
