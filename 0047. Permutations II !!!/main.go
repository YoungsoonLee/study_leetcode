package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

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
		return
	}

	used := make(map[int]bool, n-cur) // !!!

	for i := 0; i < n; i++ {

		fmt.Println("i: ", i, "cur: ", cur, "n: ", n, "nums: ", nums, "vector: ", vector, "token: ", taken, "ans: ", ans, "used: ", used)

		//if !taken[i] {
		if !taken[i] && !used[nums[i]] { // !!!
			used[nums[i]] = true // !!!

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
	a := []int{1, 1, 2}
	permuteUnique(a)
}
