package main

const base = 'z' - 'a'
const mod = int(1e9 + 7)
const maxLength = 301

func palindromePairs(words []string) [][]int {
	p := [maxLength]int{}
	p[0] = 1
	for i := 1; i < maxLength; i++ {
		p[i] = p[i-1] * base % mod
	}

	hash := make([]int, len(words))
	revHash := make([]int, len(words))

	for idx, word := range words {
		h := 0
		rh := 0
		for idx, ch := range word {
			chInt := int(ch) - 'a'
			h = (h + chInt*p[idx]%mod) % mod
			rh = (rh + chInt*p[len(word)-idx-1]%mod) % mod
		}
		hash[idx] = h
		revHash[idx] = rh
	}

	ans := [][]int{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			itoj := (hash[i] + p[len(words[i])]*hash[j]%mod) % mod
			jtoi := (revHash[j] + p[len(words[j])]*revHash[i]%mod) % mod

			if itoj == jtoi {
				ans = append(ans, []int{i, j})
			}

		}
	}

	return ans
}
