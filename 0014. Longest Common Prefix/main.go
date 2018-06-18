package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	bs := strs[getMinStringIndex(strs)]
	fmt.Println("bs: ", bs)

	bsi := len(bs)
	fmt.Println("bsi: ", bsi)

	for i := len(bs) - 1; i >= 0; i-- {
		for _, s := range strs {
			fmt.Println("rs: ", s)
			if !strings.HasPrefix(s, bs) {
				bsi--
				bs = bs[:bsi]
			}
		}
		fmt.Println("rbs: ", bs)
		fmt.Println("rbsi: ", bsi)
	}

	fmt.Println(bs[:bsi])
	return bs[:bsi]
}

func getMinStringIndex(strs []string) int {
	index := 0
	min := len(strs[0])

	for i, s := range strs {
		if min > len(s) {
			min = len(s)
			index = i
		}
	}

	return index
}

func main() {
	s := []string{"flower", "flow", "flight"}
	longestCommonPrefix(s)
}
