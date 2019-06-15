package main

import (
	"fmt"
	"sort"
)

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}

	dict := make(map[string]bool, len(wordDict))
	length := make(map[int]bool, len(wordDict))

	for _, w := range wordDict {
		length[len(w)] = true
		dict[w] = true
	}

	//fmt.Println(length)
	//fmt.Println(dict)

	sizes := make([]int, 0, len(length))
	for k := range length {
		sizes = append(sizes, k)
	}

	//fmt.Println(sizes)
	sort.Ints(sizes)
	//fmt.Println(sizes)

	// dp [i] == true, 등등. wordBreak (s [: i + 1], wordDict) == true
	dp := make([]bool, len(s)+1)
	dp[0] = true
	n := len(s)

	for i := 0; i <= n; i++ {
		fmt.Println("dp[i], i : ", dp[i], i)

		if !dp[i] {
			continue
		}

		for _, size := range sizes {
			if i+size <= n {
				dp[i+size] = dp[i+size] || dict[s[i:i+size]]
			}
		}

	}

	fmt.Println(dp)

	return dp[n]

}

func main() {
	s := "applepenapple"
	w := []string{"apple", "pen"}
	fmt.Println(wordBreak(s, w))
}
