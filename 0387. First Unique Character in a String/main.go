package main

import (
	"fmt"
)

func firstUniqChar(s string) int {

	rec := make([]int, 26)
	for _, b := range s {
		fmt.Println(b, 'a', b-'a')
		rec[b-'a']++
	}

	fmt.Println(rec)

	for i, b := range s {
		if rec[b-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}
