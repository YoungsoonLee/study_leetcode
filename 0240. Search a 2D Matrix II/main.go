package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	/* my solution
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	*/

	i, j := m-1, 0
	for 0 <= i && j < n {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			j++
			// 행렬 [i] [j]보다 작은 모든 원소 (i, j)를 제외합니다.
		} else {
			i--
			// 행렬 [i] [j]보다 큰 모든 원소 (i, j)를 제외합니다.
		}
	}
	//이 메서드의 효율성은 m 및 n의 상대 크기에 따라 다릅니다.
	// m == n 일 때 가장 효율적
	// m == 1 또는 n == 1 일 때 효율성이 가장 낮습니다.

	return false
}

func main() {

}
