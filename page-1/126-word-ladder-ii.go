package main

import (
	"math"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordList = append(wordList, beginWord)
	wordSet := make(map[string]int)
	for idx, word := range wordList {
		wordSet[word] = idx
	}

	endWordIdx, exist := wordSet[endWord]
	if !exist {
		return [][]string{}
	}

	queue := make([]int, 0)

	prev := make([][]int, len(wordList))
	for idx := range wordList {
		prev[idx] = make([]int, 0)
	}

	distance := make([]int, len(wordList))
	for i := 0; i < len(distance); i++ {
		distance[i] = math.MaxInt32
	}

	queue = append(queue, len(wordList)-1)
	distance[len(wordList)-1] = 0

	for len(queue) > 0 {
		nextQueue := make([]int, 0)
		for _, pos := range queue {
			curWord := wordList[pos]
			if curWord == endWord {
				break
			}

			for k := 0; k < len(curWord); k++ {
				for c := 'a'; c <= 'z'; c++ {
					nextWord := curWord[:k] + string(c) + curWord[k+1:]

					idx, exist := wordSet[nextWord]
					if !exist {
						continue
					}

					if distance[pos]+1 < distance[idx] {
						distance[idx] = distance[pos] + 1
						prev[idx] = []int{pos}
						nextQueue = append(nextQueue, idx)
					} else if distance[pos]+1 == distance[idx] {
						prev[idx] = append(prev[idx], pos)
					}
				}
			}
		}

		queue = nextQueue
	}

	answer := [][]string{}

	if distance[endWordIdx] == math.MaxInt32 {
		return answer
	}

	traceResult([]string{endWord}, beginWord, wordList, wordSet, prev, &answer)

	return answer
}

func traceResult(curWords []string, beginWord string, wordList []string, wordSet map[string]int, prev [][]int, answer *[][]string) {
	lastWord := curWords[len(curWords)-1]
	if lastWord == beginWord {
		*answer = append(*answer, reverse(curWords))
		return
	}
	idx := wordSet[lastWord]
	prevIndexes := prev[idx]

	for _, idx := range prevIndexes {
		traceResult(append(curWords, wordList[idx]), beginWord, wordList, wordSet, prev, answer)
	}
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
