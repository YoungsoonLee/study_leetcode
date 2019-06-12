package main

import "fmt"

func candy(ratings []int) int {
	n := len(ratings)
	if n < 1 {
		return n
	}

	// left == big than left
	left := make([]int, n)
	// right == big than right
	right := make([]int, n)

	left[0] = 1
	right[n-1] = 1

	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			// 나는 왼쪽보다 크다.
			// 그래서 그의 숫자는 왼쪽에있는 숫자보다 하나 더 많습니다.
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}

		if ratings[n-i-1] > ratings[n-i] {
			// n-i-1이 올바른 것보다 큽니다.
			// 그래서 그의 번호는 오른쪽에있는 것보다 하나 더 많습니다.
			right[n-i-1] = right[n-i] + 1
		} else {
			right[n-i-1] = 1
		}
	}

	fmt.Println(left)
	fmt.Println(right)

	res := 0
	for i := 0; i < n; i++ {
		// 실제 수는 max (left [i], right [i])
		res += max(left[i], right[i])
	}

	fmt.Println(res)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
.
/*
func candy(ratings []int) int {
	m := make(map[int]int)

	for i := 0; i < len(ratings)-1; i++ {
		c := ratings[i]
		n := ratings[i+1]
		fmt.Println(c, n)

		if c < n {
			if _, ok := m[c]; ok {
				m[c]++
			} else {
				m[c] = 2
			}
			m[n]++
		} else if n < c {
			if _, ok := m[n]; ok {
				m[n]++
			} else {
				m[n] = 2
			}
			m[c]++
		} else {
			if _, ok := m[n]; ok {
				m[n]++
			}
		}
	}

	fmt.Println(m)
	sum := 0
	for _, v := range m {
		sum += v
	}

	fmt.Println(sum)
	return sum
}
*/

func main() {
	c := []int{1, 2, 2}
	candy(c)
}
