package main

import "fmt"

func Solution(S string) int {
	// write your code in Go 1.4
	pre := make([]string, 0)
	suf := make([]string, 0)

	/*
		for i := 0; i < len(S); i++ {
			//fmt.Println(S[:i])
			pre = append(pre, S[:i])
		}

		for j := len(S) - 1; j > 0; j-- {
			//fmt.Println(S[j:])
			suf = append(suf, S[j:])
		}
	*/

	i, j := 0, len(S)-1
	for {
		if i == len(S)-1 && j == 0 {
			break
		}
		pre = append(pre, S[:i])
		suf = append(suf, S[j:])
		i++
		j--
	}

	fmt.Println(pre)
	fmt.Println(suf)

	ls := 0
	pM := make(map[string]int, len(pre))
	for _, p := range pre {
		pM[p] = len(p)
	}

	for _, s := range suf {
		if _, exists := pM[s]; exists {
			ls = max(ls, pM[s])
		}
	}
	fmt.Println(ls)
	return ls
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "abbabba"
	Solution(s)
}
