package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	n := len(nums)

	// vector is a set of possible answers
	vector := make([]int, n)

	// taken[i] == true means that nums[i] already appears in the vector
	taken := make([]bool, n)

	var ans [][]int

	makePermutation(0, n, nums, vector, taken, &ans)

	fmt.Println("ans: ", ans)
	return ans
}

// This way is compared to my original way.
// Increased the number of comparisons
// However, the number of times the slice is generated is reduced.
// So, the speed will be faster.
func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
	}

	for i := 0; i < n; i++ {

		fmt.Println("cur: ", cur, "n: ", n, "nums: ", nums, "vector: ", vector, "token: ", taken, "ans: ", ans)

		if !taken[i] {
			// Prepare to use nums[i], so take[i] == true
			taken[i] = true
			vector[cur] = nums[i]

			makePermutation(cur+1, n, nums, vector, taken, ans)

			// in the next loop
			// vector[cur] = nums[i+1]
			// So, in this loop, restore nums[i] free
			taken[i] = false
		}
	}
}

func main() {
	a := []int{1, 2, 3}
	permute(a)
}
