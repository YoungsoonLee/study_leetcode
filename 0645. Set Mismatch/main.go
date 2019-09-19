package main

import "fmt"

func findErrorNums(nums []int) []int {
	// n == len (nums)
	// 그리고 num의 수는 [1, n]이고,
	// 각 nums에서 숫자를 뒤집는다. [abs (nums [i]) - 1]
	// 작동 중에,
	// 역전되기 전에 nums [abs (nums [i]) - 1]은 이미 음수이며 abs (nums [i])가 반복 된 숫자임을 나타냅니다
	// 작업이 끝나면,
	// nums [i]> 0, nums [i]가 되돌려지지 않았 음을 나타냅니다. i + 1은 누락 된 숫자입니다.
	fmt.Println(nums)
	dup := 0
	for i := 0; i < len(nums); i++ {
		n := abs(nums[i])

		if nums[n-1] < 0 {
			dup = n
			//주의 : nums [n-1] <0이면 양수로 바꿀 수 없습니다
		} else {
			nums[n-1] = -nums[n-1]
		}
	}

	fmt.Println(nums, dup)

	mis := 0
	for i, v := range nums {
		if v > 0 {
			mis = i + 1
			break
		}
	}

	return []int{dup, mis}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	a := []int{2, 2}
	fmt.Println(findErrorNums(a))
}
