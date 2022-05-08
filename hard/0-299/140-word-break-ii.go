package main

type trieNode struct {
	nextChar map[rune]*trieNode
	isEnd    bool
}

var root trieNode

func (node *trieNode) insert(word string) {
	curNode := node
	for _, char := range word {
		if curNode.nextChar == nil {
			curNode.nextChar = make(map[rune]*trieNode)
		}
		nextNode, ok := curNode.nextChar[char]
		if !ok {
			nextNode = &trieNode{}
			curNode.nextChar[char] = nextNode
		}
		curNode = nextNode
	}
	curNode.isEnd = true
}

func trace(curNode *trieNode, s string, idx int, curWord string, finalResult *[]string) {
	if idx == len(s) {
		if !curNode.isEnd {
			return
		}
		*finalResult = append(*finalResult, curWord)
		return
	}

	char := s[idx]
	nextNode, ok := curNode.nextChar[rune(char)]
	if !ok {
		return
	}

	trace(nextNode, s, idx+1, curWord+string(char), finalResult)
	if nextNode.isEnd {
		trace(&root, s, idx+1, curWord+string(char)+string(' '), finalResult)
	}
}

func wordBreak(s string, wordDict []string) []string {
	// Should have this line,
	// since Leetcode is weird when work with global var.
	root = trieNode{}

	for _, word := range wordDict {
		root.insert(word)
	}
	result := []string{}

	trace(&root, s, 0, "", &result)

	return result
}
