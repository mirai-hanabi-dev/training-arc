package main

import (
	"fmt"
	"math"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordList = append(wordList, beginWord)
	wordSet := make(map[string]int)
	for idx, word := range wordList {
		wordSet[word] = idx
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

	answer := [][]string{{endWord}}

	return answer
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println(findLadders(beginWord, endWord, wordList))
}
