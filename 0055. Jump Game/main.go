package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}

		j := i - 1
		for ; j >= 0; j-- {
			// j 위치에서 0 요소를 건너 뛸 수 있습니다.
			if i-j < nums[j] {
				i = j
				break
			}
		}

		if j == -1 { // 0 요소 앞에 0을 놓을 수있는 위치가 없습니다.
			return false
		}
	}

	return true
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(canJump(a))
	//b := []int{3, 2, 1, 0, 4}
	//fmt.Println(canJump(b))
}
