package main

import "fmt"

// 문제를 이해를 못 함.
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	// init dp
	// dp [i] [j] [i + j] == true는 s1 [: i]와 s2 [: j]가 s3 [: i + j]를 생성 할 수 있다는 것을 의미합니다.
	dp := make([][][]bool, m+1)
	fmt.Println("dp1: ", dp)

	for i := 0; i <= m; i++ {
		dp[i] = make([][]bool, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]bool, m+n+1)
		}
	}
	fmt.Println("dp2: ", dp)
	dp[0][0][0] = true

	//
	for i := 1; i <= m; i++ {
		dp[i][0][i] = dp[i-1][0][i-1] && s1[i-1] == s3[i-1]
	}

	fmt.Println("dp3: ", dp)

	return true
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	isInterleave(s1, s2, s3)
}
