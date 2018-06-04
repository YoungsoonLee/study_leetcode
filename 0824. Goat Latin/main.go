package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {
	ss := strings.Split(S, " ")

	for i := range ss {
		ss[i] = handleWord(ss[i], i)
	}

	return strings.Join(ss, " ")
}

func handleWord(s string, i int) string {
	postfix := "ma" + strings.Repeat("a", i+1)

	if isBeginWithVowel(s) {
		return s + postfix
	}

	return s[1:] + s[0:1] + postfix
}

func isBeginWithVowel(s string) bool {
	switch s[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

/* below is my solution
const v = "aeiouAEIOU"
func toGoatLatin(S string) string {
	ret := make([]string, 0)
	for _, w := range strings.Split(S, " ") {
		fmt.Println("w: ", w)
		checkFirstVowel(w, &ret)
		checkConsonant(w, &ret)
	}

	addA(&ret)

	fmt.Println(strings.Join(ret, " "))
	return strings.Join(ret, " ")
}

func checkFirstVowel(w string, ret *[]string) {
	//v := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	if strings.Contains(v, string(w[0])) {
		*ret = append(*ret, string(w)+"ma")
	}
}

func checkConsonant(w string, ret *[]string) {
	if !strings.Contains(v, string(w[0])) {
		w = w[1:] + string(w[0]) + "ma"
		*ret = append(*ret, w)
	}
}

func addA(ret *[]string) {
	for i, w := range *ret {
		fmt.Println(w)
		(*ret)[i] = w + geta(i)
	}
}

func geta(n int) string {
	// strings.Repeat("a", i+1) !!!!!
	s := "a"
	for i := 1; i <= n; i++ {
		s += "a"
	}
	return s
}
*/

func main() {
	a := "I speak Goat Latin"
	fmt.Println(toGoatLatin(a))
}
