package main

func isToeplitzMatrix(matrix [][]int) bool {
	r, c := len(matrix), len(matrix[0])

	for i := 0; i+1 < r; i++ {
		for j := 0; j+1 < c; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}

	return true
}

func main() {

}
