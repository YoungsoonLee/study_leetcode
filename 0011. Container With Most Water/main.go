package main

import "fmt"

func maxArea(height []int) int {
	// 양쪽 끝에서 보면서 적어도 너비가 최대가되도록하십시오.
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)

		area := h * (j - i)
		if max < area {
			max = area
		}

		// 영역쪽으로 더 커질 가능성이있는 방향으로 변경합니다.
		if a < b {
			i++
		} else {
			j--
		}
	}

	fmt.Println(max)
	return max
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	a := []int{1, 1, 9, 8, 5, 7}
	maxArea(a)
}
