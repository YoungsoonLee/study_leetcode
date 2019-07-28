package main

import (
	"fmt"
)

/**
 * my solution - fail
 *
func isAnagram(s string, t string) bool {
	sm := make(map[string]int, len(s))
	tm := make(map[string]int, len(t))

	if len(s) == 0 && len(s) == 0 {
		return true
	}

	if len(s) != len(t) {
		return false
	}

	s = strings.ToLower(s)
	t = strings.ToLower(t)

	for _, c := range s {
		if _, ok := sm[string(c)]; ok {
			sm[string(c)]++
		}
		sm[string(c)] = 1
	}

	for _, c := range t {
		if _, ok := tm[string(c)]; ok {
			tm[string(c)]++
		}
		tm[string(c)] = 1
	}

	fmt.Println(sm, tm)

	for _, tv := range t {
		if tm[string(tv)] == sm[string(tv)] {
			return true
		} else {
			return false
		}
	}

	return false
}
*/

/*
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sm := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sm[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if _, ok := sm[t[i]]; ok {
			sm[t[i]]--
		}
	}

	for _, v := range sm {
		if v > 0 {
			return false
		}
	}

	return true
}
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)

	fmt.Println(sr, tr)

	rec := make(map[rune]int, len(sr))
	//fmt.Println(rec)

	for i := range sr {
		rec[sr[i]]++
		rec[tr[i]]--
	}

	fmt.Println(rec)

	for _, n := range rec {
		//fmt.Println("n: ", n)
		if n != 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "dog"
	s2 := "gdo"

	fmt.Println(isAnagram(s1, s2))
}
