package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	visited := make(map[string]bool)

	step := 0

	queue := make([]string, 0)
	queue = append(queue, beginWord)
	visited[beginWord] = true

	for len(queue) > 0 {
		step += 1
		nextQueue := make([]string, 0)
		for _, curWord := range queue {
			if curWord == endWord {
				return step
			}

			for k := 0; k < len(curWord); k++ {
				for i := 'a'; i <= 'z'; i++ {
					nextWord := curWord[:k] + string(i) + curWord[k+1:]
					if _, exist := visited[nextWord]; exist {
						continue
					}
					if _, exist := wordSet[nextWord]; !exist {
						continue
					}
					nextQueue = append(nextQueue, nextWord)
					visited[nextWord] = true
				}
			}
		}

		queue = nextQueue
	}

	return 0
}
