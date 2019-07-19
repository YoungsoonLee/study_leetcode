package main

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	maxEdge := 0
	// dp [i] [j] == 점 (i, j)의 오른쪽 아래 모서리에서 가장 큰 사각형의 길이
	// TODO : dp [i] [j]는 위쪽, 왼쪽 위 및 왼쪽 데이터와 만 관련이 있으므로
	// dp는 1 차원으로 압축 될 수 있습니다
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxEdge = 1
		}
	}

	for j := 1; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxEdge = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
				maxEdge = max(maxEdge, dp[i][j])
			}
		}
	}

	return maxEdge * maxEdge
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
