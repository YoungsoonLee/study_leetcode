package main

import (
	"fmt"
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	fmt.Println(words)

	m := make(map[string]bool, len(words))
	res := words[0]

	for _, w := range words {
		n := len(w)
		fmt.Println(n, w, m[w], w[:n-1], m[w[:n-1]])

		if n == 1 {
			m[w] = true
		} else if m[w[:n-1]] {
			m[w] = true
			if len(res) < len(w) {
				res = w
			}
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	w := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	longestWord(w)
}
