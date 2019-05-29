package main

import "fmt"

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)

	dp := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = 1
	}

	fmt.Println("init dp: ", dp, m, n)

	// dp [i] 업데이트 전의 값, 즉 s [: i]와 t [: j] 사이의 일치
	// s [: i] 및 t [: j + 1]의 일치하는 값인 dp [i + 1]을 업데이트 할 때 사용될 수 있습니다.
	// prev로 저장
	var prev int

	for j := 0; j < n; j++ {
		dp[j], prev = 0, dp[j]
		fmt.Println("mid dp: ", dp, m, n, prev)

		for i := j + 1; i < m+1; i++ {
			// for dp [i]
			// s [: i]의 정규화 된 부분 문자열은 s [i-1]이 포함되어 있는지 여부에 따라 두 부분으로 나눌 수 있습니다.
			// 파트 1 :
			// 		** not **는 s [i-1]을 포함한다.
			//		이 부분의 번호는 dp [i-1]과 같습니다.
			// 파트 II :
			// 		s [i-1]을 포함합니다.
			//		이 부분은 s [i-1] == t [j-1] 일 때만 존재합니다.
			//		있는 경우이 부분의 수는 prev와 같습니다.
			if t[j] == s[i-1] {
				dp[i], prev = dp[i-1]+prev, dp[i]
			} else {
				dp[i], prev = dp[i-1], dp[i]
			}
		}
	}

	fmt.Println("last dp: ", dp, m, n)
	return dp[m]
}

func main() {
	s := "rabbbit"
	t := "rabbit"
	numDistinct(s, t)
}
