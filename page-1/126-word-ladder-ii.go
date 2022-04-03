package main

import (
	"fmt"
)

var cache = make(map[string]bool)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	visited := make(map[string]bool)
	visited[beginWord] = true

	return [][]string{}
}

func ableTransform(first string, second string) bool {
	if first > second {
		return ableTransform(second, first)
	}
	concat := first + "-" + second
	if v, ok := cache[concat]; ok {
		return v
	}

	cache[concat] = true

	if len(first) != len(second) {
		cache[concat] = false
		return false
	}

	diffFlag := false
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			if diffFlag {
				cache[concat] = false
				return false
			}
			diffFlag = true
		}
	}

	return cache[concat]
}

func main() {
	beginWord := "hit"
	endWord := "hot"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println(findLadders(beginWord, endWord, wordList))
	fmt.Println(ableTransform(beginWord, endWord))
}
