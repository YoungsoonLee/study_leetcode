package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	fmt.Println(words)

	for i, w := range words {
		words[i] = reverse(w)
	}

	fmt.Println(words)
	fmt.Println(strings.Join(words, " "))
	return strings.Join(words, " ")
}

/*
func reverse(s string) string {
	bytes := []byte(s) // !!!!!
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
*/

/*
func reverseWords(s string) string {
	ret := ""
	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ret += reverse(s[idx:i]) + " "
			idx = i + 1
		}
	}

	ret += " " + reverse(s[idx:])
	fmt.Println(ret)
	return ret
}
*/

func reverse(s string) string {

	bytes := []byte(s) // !!!!!
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)

}

func main() {
	a := "Let's take LeetCode contest"
	reverseWords(a)
}
