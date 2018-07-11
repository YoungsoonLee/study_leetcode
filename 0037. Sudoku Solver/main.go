package main

func solveSudoku(board [][]byte) {
	if !fill(board, '1', 0) {
		panic("No solution to this question")
	}
}

func fill(board [][]byte, n byte, block int) bool {
	if block == 9 {
		// All the blocks are filled and successfully found
		return true
	}

	if n == '9'+1 {
		// block has been filled, fill in block+1
		return fill(board, '1', block+1)
	}

	rowBegin := (block / 3) * 3
	colBegin := (block % 3) * 3

	for r := rowBegin; r < rowBegin+3; r++ {
		for c := colBegin; c < colBegin+3; c++ {
			if board[r][c] == n {
				// There is already n in the block, no need to fill in
				// to fill in n+1
				return fill(board, n+1, block)
			}
		}
	}

	// check if (r, c) can be stored n
	// Use anonymous functions to avoid passing parameters
	isAvaliable := func(r, c int) bool {
		// The character at the current position needs to be '.'
		if board[r][c] != '.' {
			return false
		}

		// (r,c) can't have rows and columns
		// It can be reflected here, and the advantage of filling in the block.
		for i := 0; i < 9; i++ {
			if board[r][i] == n || board[i][c] == n {
				return false
			}
		}

		return true
	}

	for r := rowBegin; r < rowBegin+3; r++ {
		for c := colBegin; c < colBegin+3; c++ {
			if isAvaliable(r, c) {
				board[r][c] = n
				if fill(board, n+1, block) {
					return true
				}

				// Restore (r,c) to move n to the next possible location later
				board[r][c] = '.'

				// print(board, n, block)
			}
		}
	}

	// n Nowhere to put in this block.
	// return false to let n-1 adjust the position.
	return false
}

func main() {

}
