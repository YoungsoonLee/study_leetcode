package main

import "fmt"

// 리버스의 재귀 이용
func rotate(nums []int, k int) {
	n := len(nums)

	if k > n {
		k %= n
	}

	fmt.Println(n, k)

	if k == 0 || k == n {
		return
	}

	//nums = append(nums[:k-1], nums[:k+1]...)
	//fmt.Println(nums)

	reverse(nums, 0, n-1) // all reverse !!!!!!
	//fmt.Println(nums)
	reverse(nums, 0, k-1) // reverse
	//fmt.Println(nums)
	reverse(nums, k, n-1) // reverse
	fmt.Println(nums)

}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(a, k)
}
