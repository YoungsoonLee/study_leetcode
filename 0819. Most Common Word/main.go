package main

import (
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	isBanned := make(map[string]bool)
	for _, b := range banned {
		isBanned[b] = true
	}

	ss := split(paragraph)

	count := make(map[string]int, len(ss))

	for _, s := range ss {
		if isBanned[s] {
			continue
		}
		count[s]++
	}

	res := ""
	max := 0

	for s, c := range count {
		if max < c {
			max = c
			res = s
		}
	}

	return res
}

func split(s string) []string {
	p := replace(s)
	p = strings.ToLower(p)
	return strings.Split(p, " ")
}

func replace(s string) string {
	chars := []string{"! ", "? ", ", ", "' ", "; ", ". ", "."}
	for _, c := range chars {
		s = strings.Replace(s, c, " ", -1)
	}
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	p := "Bob hit a ball, the hit BALL flew far after it was hit."
	b := []string{"hit"}
	mostCommonWord(p, b)
}
