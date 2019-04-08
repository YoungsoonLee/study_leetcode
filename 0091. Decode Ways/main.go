package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	memo := make([]int, len(s))

	// init memo
	for i := range memo {
		memo[i] = -1
	}

	result := track(s, 0, memo)
	fmt.Println(result)
	return result
}

func track(s string, start int, memo []int) int {
	if start >= len(s) {
		return 1
	}

	if s[start] == '0' {
		return 0
	}

	if memo[start] > -1 {
		return memo[start]
	}

	count := 0
	count += track(s, start+1, memo)
	if start+1 < len(s) {
		num, err := strconv.Atoi(s[start : start+2])
		if err == nil && num <= 26 && num > 1 {
			count += track(s, start+2, memo)
		}
	}
	memo[start] = count
	return count
}

func main() {
	s := "226"
	numDecodings(s)
}
