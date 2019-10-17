package main

import (
	"fmt"
	"strings"
)

/*
func letterCasePermutation(s string) []string {
	fmt.Println("s: ", s) // !!!!!
	size := len(s)
	if size == 0 {
		return []string{""}
	}

	lastByte := s[size-1]
	postfixs := make([]string, 1, 2)
	//fmt.Println("init postfixs: ", postfixs)

	postfixs[0] = string(lastByte)
	fmt.Println("postfixs: ", postfixs)

	if b, ok := check(lastByte); ok {
		postfixs = append(postfixs, string(b))
	}

	// 재귀를 사용하여 접두사를 계산합니다.
	prefixs := letterCasePermutation(s[:size-1]) // !!!! 재귀
	fmt.Println("prefixs: ", prefixs)

	// res는 접두사와 접미사의 곱입니다.
	res := make([]string, 0, len(prefixs)*len(postfixs))

	for _, pre := range prefixs {
		for _, post := range postfixs {
			fmt.Println("pre, post", pre, post)
			res = append(res, pre+post)
		}
	}

	fmt.Println(res)
	return res
}
*/

func letterCasePermutation(S string) []string {

	// get string count first
	count := 0
	m := make(map[int]string)

	for i, c := range S {
		//fmt.Println(c)

		if isChar(c) {
			count++
			m[i] = string(c)
		}
	}

	res := make([]string, count*2)

	//fmt.Println(res)

	for i := 0; i < count*2; i++ {
		for j := 0; j < len(S)-1; j++ {
			if _, ok := m[j]; ok {
				// fmt.Println(j)
				// change(S[j]) + S[j+1:]
				fmt.Println(fmt.Sprintf("%s", S[j+1:]))

				res[i] = change(S[j]) + S[j+1:]
			}
		}
	}

	//fmt.Println(res)

	return res

}

func isChar(c rune) bool {
	// fmt.Println(c)
	// fmt.Println('Z')
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func change(s byte) string {
	if s >= 'a' && s <= 'z' {
		return strings.ToUpper(string(s))
	}
	return strings.ToLower(string(s))
}

func check(b byte) (byte, bool) {
	if 'a' <= b && b <= 'z' {
		return b + 'A' - 'a', true
	}
	if 'A' <= b && b <= 'Z' {
		return b + 'a' - 'A', true
	}
	return 0, false
}

func main() {
	s := "a1b2"
	letterCasePermutation(s)
}
