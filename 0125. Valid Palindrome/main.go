package main

import (
	"fmt"
	"strings"
)

// rotor = true

// 하나의 문자로 만들수 있는 여러개의 단어나 스트링
// my solution - failed
/*
func isPalindrome(s string) bool {
	//s = strings.ToLower(strings.Replace(s, " ", "", -1))
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return isPalindrome(s[1 : len(s)-1])
}
*/

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}

	return false
}

// recursive
func isPalindrome3(s string) bool {
	//fmt.Println(s)
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return isPalindrome3(s[1 : len(s)-1])
}

func main() {
	s := "rotor"
	fmt.Println(isPalindrome3(s))
}
