package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	r := len(matrix)
	if r == 0 {
		return false
	}

	c := len(matrix[0])
	if c == 0 {
		return false
	}

	if target < matrix[0][0] || matrix[r-1][c-1] < target {
		return false
	}

	// find row !!!!
	row := 0
	for row < r && matrix[row][0] <= target {
		row++
	}
	row--

	fmt.Println("row: ", row)

	i, j := 0, c-1
	for i <= j {
		med := (i + j) / 2
		switch {
		case matrix[row][med] < target:
			i = med + 1
		case matrix[row][med] > target:
			j = med - 1
		default:
			return true
		}
	}

	return false

	/* my solution is bad
	a := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			a = append(a, matrix[i][j])
		}
	}

	fmt.Println("a: ", a)

	// binary search
	low, high := 0, len(a)
	var mid int

	for low <= high {
		mid = (low + high) / 2

		switch {
		case a[mid] < target:
			low = mid + 1
		case a[mid] > target:
			high = mid - 1
		default:
			return true
		}
	}

	return false
	*/
}

func main() {
	n := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	t := 30
	fmt.Println(searchMatrix(n, t))
}
