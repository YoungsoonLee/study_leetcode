package main

var m map[string]int

func longestPalindromeSubseq(s string) int {

	if m == nil {
		m = make(map[string]int)
	}

	if len(s) <= 1 {
		return len(s)
	}

	cnt := 0
	if cnt, ok := m[s]; ok {
		return cnt
	}

	if s[0] == s[len(s)-1] {
		//return  2 + longestPalindromeSubseq(s[1:len(s)-1])
		cnt = 2 + longestPalindromeSubseq(s[1:len(s)-1])

	} else {
		l := longestPalindromeSubseq(s[1:])
		r := longestPalindromeSubseq(s[:len(s)-1])
		//return max(l, r)
		if l > r {
			cnt = l
		} else {
			cnt = r
		}
	}
	m[s] = cnt
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
