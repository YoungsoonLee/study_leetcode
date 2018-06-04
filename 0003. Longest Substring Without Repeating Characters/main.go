package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0

	nm := make(map[string]int, 0) // current index of character
	for i, j := 0, 0; j < n; j++ {
		if _, ok := nm[string(s[j])]; ok {
			i = max(nm[string(s[j])], i)
			fmt.Println("i: ", nm[string(s[j])], i)
		}
		ans = max(ans, j-i+1) // !!!!!
		nm[string(s[j])] = j + 1

		fmt.Println(nm, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func lengthOfLongestSubstring(s string) int {
	// location [s [i]] == j는 다음을 의미합니다.
	// s의 i 번째 문자열은 s의 j 위치에 마지막으로 표시되므로 s [j + 1 : i]에는 s [i]가 없습니다.
	// location [s [i]] == -1은 다음을 의미합니다. s [i]가 s에 처음 나타납니다.

	location := [256]int{} // 입력 문자열에 ASCII 문자 만 있다고 가정하기 때문에 // 길이가 256입니다.
	for i := range location {
		location[i] = -1 // 이전에 보지 않은 모든 문자를 설정합니다.
	}

	fmt.Println("init location: ", location)

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 설명 s [i]가 s [left : i + 1]에서 반복되었습니다.
		// 그리고 s [i]의 마지막 발생은 위치 [s [i]]에있었습니다.
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // s [i] 문자와 앞 부분을 s [left : i + 1]에서 제거합니다.
		} else if i+1-left > maxLen {
			fmt.Println(s[left : i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}
	fmt.Println(location)
	fmt.Println(maxLen)

	return maxLen
}
*/

func main() {
	a := "pwwkew"
	lengthOfLongestSubstring(a)
}
