package main

import "fmt"

func sortColors(nums []int) {
	i, j, k := 0, 0, len(nums)-1

	// 루프의 경우 nums [i : j]는 항상 1입니다.
	// 루프가 끝나면,
	// nums [: i]는 모두 0입니다.
	// nums [j :]는 모두 2입니다.
	for j <= k {
		switch nums[j] {
		case 0:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1:
			j++
		case 2:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
		//fmt.Println(nums)
	}

	fmt.Println(nums)
}

func main() {
	n := []int{2, 0, 2, 1, 1, 0}
	sortColors(n)
}
