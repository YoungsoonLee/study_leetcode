package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix) // row
	if m == 0 {
		return 0
	}

	n := len(matrix[0]) // col
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0') //change first column
		for i := 1; i < m; i++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}

	fmt.Println("first dp: ", dp)

	max := 0
	for i := 0; i < m; i++ {
		tmp := largestRectangleArea(dp[i])
		if max < tmp {
			max = tmp
		}
	}

	return max
}

func largestRectangleArea(heights []int) int {
	// h의 끝에 -1을 추가하여 for 루프의 영역을 논리적으로 정렬합니다.
	h := append(heights, -1)
	n := len(h)

	// var maxArea, height, left, right, area int
	// 높이 배열이 오름차순으로 알려진 경우 어떻게해야합니까?
	// 예를 들어 1, 2, 5, 7, 8
	// (5 * 3) vs (7 * 2) vs (2 * 4) vs (8 * 1)
	// 또한 max (height [i] * (size-i))입니다.
	// 스택을 사용하는 목적은 이러한 오름차순 시퀀스를 작성하고 위의 방법으로 해결하는 것입니다.

	return 0
}

func main() {
	m := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}
	maximalRectangle(m)
}
