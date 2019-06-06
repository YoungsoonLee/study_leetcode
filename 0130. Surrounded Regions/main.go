package main

import "fmt"

func solve(board [][]byte) {
	m := len(board)
	if m <= 2 {
		return
	}

	n := len(board[0])
	if n <= 2 {
		return
	}

	// isEscaped [i] [j] == true, (i, j) 점이 Unicom임을 나타냅니다.
	// 또한 점 (i, j)가 검사되었는지 확인하는 데 사용할 수 있습니다.
	isEscaped := make([][]bool, m)
	for i := 0; i < m; i++ {
		isEscaped[i] = make([]bool, n)
	}

	// idxM, idxN은 각각 점 (i, j)의 좌표 값을 기록합니다.
	idxM := make([]int, 0, m*n)
	idxN := make([]int, 0, m*n)

	bfs := func(i, j int) {
		isEscaped[i][j] = true
		idxM = append(idxM, i)
		idxN = append(idxN, j)

		for len(idxM) > 0 {
			// (i, j)의 왼쪽 위, 오른쪽, 오른쪽 포인트를 차례대로 수행하고, bfs check
			i := idxM[0]
			j := idxN[0]
			idxM = idxM[1:]
			idxN = idxN[1:]

			if 0 <= i-1 && board[i-1][j] == 'O' && !isEscaped[i-1][j] {
				idxM = append(idxM, i-1)
				idxN = append(idxN, j)
				isEscaped[i-1][j] = true
			}

			if 0 <= j-1 && board[i][j-1] == 'O' && !isEscaped[i][j-1] {
				idxM = append(idxM, i)
				idxN = append(idxN, j-1)
				isEscaped[i][j-1] = true
			}

			if i+1 < m && board[i+1][j] == 'O' && !isEscaped[i+1][j] {
				idxM = append(idxM, i+1)
				idxN = append(idxN, j)
				isEscaped[i+1][j] = true
			}

			if j+1 < n && board[i][j+1] == 'O' && !isEscaped[i][j+1] {
				idxM = append(idxM, i)
				idxN = append(idxN, j+1)
				isEscaped[i][j+1] = true
			}

		}

	}

	// Unicom의 영역은 확실히 도달 할 것이므로 4 주에서 모든 Unicom 영역을 확인하십시오.
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' && !isEscaped[0][j] {
			bfs(0, j)
		}
		if board[m-1][j] == 'O' && !isEscaped[m-1][j] {
			bfs(m-1, j)
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' && !isEscaped[i][0] {
			bfs(i, 0)
		}
		if board[i][n-1] == 'O' && !isEscaped[i][n-1] {
			bfs(i, n-1)
		}
	}

	fmt.Println(isEscaped)

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			// 修改内部被占领的 'O'
			if board[i][j] == 'O' && !isEscaped[i][j] {
				board[i][j] = 'X'
			}
		}
	}

	fmt.Println(isEscaped)

	/*
		r := len(board)
		c := len(board[0])

		for i := 0; i < r-1; i++ {
			for j := 0; j < c-1; j++ {
				if i == 0 || i == r-1 {
					// do nothing
				}
				if j == 0 || j == c-1 {
					// do nothing
				}

				if board[i][j] == 'O' {
					board[i][j] = 'X'
				}
			}
		}
	*/

}

func main() {

}
