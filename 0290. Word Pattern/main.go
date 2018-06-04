package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	hm := make(map[string]string, len(pattern))
	ss := strings.Split(str, " ")
	if len(pattern) != len(ss) {
		return false
	}
	//
	for i := 0; i < len(pattern); i++ {
		if _, ok := hm[string(pattern[i])]; ok {
			if hm[string(pattern[i])] != string(ss[i]) {
				fmt.Println(hm)
				return false
			}
		}
		hm[string(pattern[i])] = string(ss[i])
	}

	//var preV string
	prev := ""
	for _, v := range hm {
		if v == prev {
			return false
		}
		prev = v
	}
	return true
}

func wordPattern2(pattern string, str string) bool {
	ps := strings.Split(pattern, "")
	ss := strings.Split(str, " ")

	if len(ps) != len(ss) {
		return false
	}

	return isMatch(ps, ss) && isMatch(ss, ps)
}

func isMatch(s1, s2 []string) bool {
	size := len(s1)

	m := make(map[string]string, size)

	var i int
	var w string
	var ok bool

	for i = 0; i < size; i++ {
		if w, ok = m[s1[i]]; ok {
			if w != s2[i] {
				return false
			}
		} else {
			m[s1[i]] = s2[i]
		}
	}
	return true
}

func main() {
	p := "abba"
	str := "dog dog dog dog"
	fmt.Println(wordPattern(p, str))
}
