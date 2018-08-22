package main

import "fmt"

func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)

	dp := [][]bool{}

	for i := 0; i < ls+1; i++ {
		rt := make([]bool, lp+1)
		dp = append(dp, rt)
	}
	fmt.Println("init dp: ", dp)

	// dp [i] [j] == true는 s [: i + 1]가 p [: j + 1]과 일치 할 수 있음을 의미합니다.
	dp[0][0] = true

	for j := 1; j <= lp; j++ {
		if p[j-1] == '*' {
			// p [j-1] == '*'일 때
			// dp [0] [j]는 이전 경기와 일치합니다.
			// p [j-1]! = '*'이면, 다음의 dp [0] [j]는 모두 false입니다.
			dp[0][j] = dp[0][j-1]
		}
	}

	fmt.Println("after first for loop dp: ", dp)

	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if p[j-1] != '*' {
				// p [j-1]! = '*'일 때
				// 단일 문자가 일치하고 이전 문자열도 일치합니다.
				dp[i][j] = (p[j-1] == s[i-1] || p[j-1] == '?') && dp[i-1][j-1]
			} else {
				// p [j-1] == '*'일 때
				// dp [i-1] [j] == true는
				// s [: i]가 p [: j + 1]과 p [j] == '*'와 일치 할 때,
				// s [: i] s [: i + 1] 뒤에 p [: j + 1]와 여전히 일치하는 문자열이옵니다.
				// dp [i] [j-1] == true는
				// s [: i + 1]이 p [: j]와 일치 할 때 p [: j], s [: i + 1] 및 p [: j + 1] 다음에 '*'를 추가하십시오.
				// 왜냐하면 '*'는 null 문자와 일치 할 수 있기 때문입니다.
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	fmt.Println("after second for loop dp: ", dp)
	fmt.Println("result dp: ", dp[ls][lp])

	return dp[ls][lp]
}

func main() {
	s := "acdcb"
	p := "a*c?b"
	isMatch(s, p)
}

/*
func isMatch(s string, p string) bool {
	//
	if len(p) == 1 && p == "*" {
		return true
	}

	if len(p) != len(s) {
		return false
	}

	for _, c := range s {
		for _, ep := range p {
			if string(ep) == "*" {
				continue
			}

			if string(c) != string(ep) {
				return false
			}
		}
	}

	return true
}
*/
