// It took me 3 tries to Accept this problem...

package main

func longestCommonPrefix(strs []string) string {
	maxRight := 0
	shoudStop := false
	base := strs[0]
	for k := 0; k < len(base)+1; k++ {
		if shoudStop {
			break
		}
		maxRight = k
		for _, str := range strs {
			if k == len(str) {
				shoudStop = true
				break
			}
			if str[k] == base[k] {
				continue
			} else {
				shoudStop = true
				break
			}
		}
	}

	return base[:maxRight]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
