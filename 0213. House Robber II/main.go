package main

func rob(nums []int) int {
	size := len(nums)

	switch size {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	return max(robbing(nums[1:]), robbing(nums[:size-1]))
}

func robbing(nums []int) int {
	// 짝수의 최대 값 레코드
	// b 홀수 비트의 최대 값을 기록합니다.
	var a, b int
	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(a, b+v)
		}
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := []int{1, 2, 3, 1}
	rob(n)
}
