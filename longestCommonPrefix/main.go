package main

import (
	"fmt"
	"strings"
)

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

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	//sort.Sort(sort.StringSlice(strs))

	bs := strs[getMinStringIndex(strs)]
	bsi := len(bs)
	//pb := false

	//fmt.Println("bsi: ", bsi)
	//fmt.Println(bs)

	for i := len(bs) - 1; i >= 0; i-- {
		for _, s := range strs {
			//if !strings.Contains(s, bs) {
			if !strings.HasPrefix(s, bs) {
				bsi--
				bs = bs[:bsi]
			}
		}
	}

	//fmt.Println("bsi: ", bsi)
	fmt.Println(bs[:bsi])
	return bs[:bsi]
}

func main() {
	a := []string{"ca", "a"}
	fmt.Println(longestCommonPrefix(a))

	s := []string{"aa", "ab"}
	fmt.Println(longestCommonPrefix(s))

	b := []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	fmt.Println(longestCommonPrefix(b))

	c := []string{"aaa", "aa", "aaa"}
	fmt.Println(longestCommonPrefix(c))

	d := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(d))
}
