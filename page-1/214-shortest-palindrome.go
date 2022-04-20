package main

var mod = int(1e9 + 7)
var base = int('z' - 'a')

func shortestPalindrome(s string) string {
	ltr := 0
	for i := range s {
		c := int(s[i] - 'a')
		ltr = (c*pow(base, i) + ltr) % mod
	}

	rtl := 0
	for i := range s {
		c := int(s[len(s)-i-1] - 'a')
		rtl = (c*pow(base, i) + rtl) % mod
	}

	addS := ""
	curltr := 0
	currtl := 0

	for i := range s {
		tmpltr := (((ltr * pow(base, len(addS))) % mod) + curltr) % mod
		tmprtl := (((currtl * pow(base, len(s))) % mod) + rtl) % mod

		if tmpltr == tmprtl {
			return addS + s
		}

		r := s[len(s)-i-1]
		c := int(r - 'a')
		curltr = (c*pow(base, len(addS)) + curltr) % mod
		currtl = (currtl*base + c) % mod

		addS += string(r)
	}
	return ""
}

func pow(a, x int) int {
	if x == 0 {
		return 1
	}
	p := pow(a, x/2)
	m := (p * p) % mod
	if x%2 == 0 {
		return m
	}
	return (m * a) % mod
}
