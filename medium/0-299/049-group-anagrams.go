package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := map[string][]string{}

	for _, str := range strs {
		count := map[rune]int{}
		for _, c := range str {
			count[c]++
		}

		keys := []rune{}
		for key := range count {
			keys = append(keys, key)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})

		sortStr := ""
		for _, key := range keys {
			sortStr += strings.Repeat(string(key), count[key])
		}

		if _, ok := res[sortStr]; !ok {
			res[sortStr] = []string{}
		}
		res[sortStr] = append(res[sortStr], str)
	}

	values := [][]string{}
	for key := range res {
		values = append(values, res[key])
	}
	return values
}
