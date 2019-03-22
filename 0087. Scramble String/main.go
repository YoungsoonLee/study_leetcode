package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	n := len(s1)

	// 두 문자열이 같은 문자를 가지고 있는지 확인
	rec := make([]int, 256)
	for i := 0; i < n; i++ {
		rec[s1[i]]++
		rec[s2[i]]--
	}
	for i := range rec {
		if rec[i] != 0 {
			return false // 완전히 다른 문자열.
		}
	}

	// 하위 문자열이 스크램블인지 확인
	for i := 1; i < n; i++ {
		// 두 문자열 사이에 점이 있으면 왼쪽의 하위 문자열이 스크램블 문자열이고 오른쪽의 하위 문자열도 스크램블 문자열이며 true이면
		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}

		// 두 문자열의 중간에 점이 있으면 s1의 왼쪽에있는 부분 문자열과 s2의 오른쪽에있는 부분 문자열이 스크램블 문자열이며
		// s1의 오른쪽에있는 부분 문자열과 s2의 왼쪽에있는 부분 문자열도 스크램블 문자열이며 true입니다.
		if isScramble(s1[0:i], s2[n-i:]) && isScramble(s1[i:], s2[0:n-i]) {
			return true
		}
	}

	return false
}

func main() {
	s1 := "great"
	s2 := "rgeat"
	fmt.Println(isScramble(s1, s2))
}
