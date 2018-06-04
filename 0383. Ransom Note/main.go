package main

import (
	"fmt"
)

/**
 * my solution - failed
func canConstruct(ransomNote string, magazine string) bool {
	return strings.Contains(magazine, ransomNote)
}
*/
func canConstruct(ransomNote string, magazine string) bool {
	mc := getCount(magazine)
	//fmt.Println(mc)

	for _, b := range ransomNote {
		mc[b-'a']--
		if mc[b-'a'] < 0 {
			fmt.Println(mc)
			return false
		}
	}
	fmt.Println(mc)
	return true
}

func getCount(s string) []int {
	res := make([]int, 26)
	for i := range s {
		fmt.Println(i, s[i], 'a', s[i]-'a')
		res[s[i]-'a']++
	}

	fmt.Println(res)
	return res
}

func main() {
	r := "a"
	m := "b"
	fmt.Println(canConstruct(r, m))
}

//"fffbfg"
//"effjfggbffjdgbjjhhdegh"
