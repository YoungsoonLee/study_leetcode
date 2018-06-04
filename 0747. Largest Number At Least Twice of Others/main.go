package main

import (
	"fmt"
)

func dominantIndex(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	// i1은 a에서 첫 번째로 큰 숫자의 색인 번호입니다.
	// i2는 a에서 두 번째로 큰 숫자의 색인 번호입니다.
	i1, i2 := 0, 1
	if nums[i1] < nums[i2] {
		i1, i2 = i2, i1
	}

	// 실제 i1과 i2를 찾습니다.
	for i := 2; i < n; i++ {
		if nums[i1] < nums[i] {
			i2, i1 = i1, i
		} else if nums[i2] < nums[i] {
			i2 = i
		}
	}

	if nums[i2] == 0 || nums[i1]/nums[i2] >= 2 {
		return i1
	}
	return -1
}

func main() {
	n := []int{1, 2, 3, 4}
	fmt.Println(dominantIndex(n))
}
