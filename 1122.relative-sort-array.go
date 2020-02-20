import "fmt"

/*
 * @lc app=leetcode id=1122 lang=golang
 *
 * [1122] Relative Sort Array
 *
 * https://leetcode.com/problems/relative-sort-array/description/
 *
 * algorithms
 * Easy (66.97%)
 * Likes:    417
 * Dislikes: 30
 * Total Accepted:    40.3K
 * Total Submissions: 59.8K
 * Testcase Example:  '[2,3,1,3,2,4,6,7,9,2,19]\n[2,1,4,3,9,6]'
 *
 * Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all
 * elements in arr2 are also in arr1.
 *
 * Sort the elements of arr1 such that the relative ordering of items in arr1
 * are the same as in arr2.  Elements that don't appear in arr2 should be
 * placed at the end of arr1 in ascending order.
 *
 *
 * Example 1:
 * Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
 * Output: [2,2,2,1,4,3,3,9,6,7,19]
 *
 *
 * Constraints:
 *
 *
 * arr1.length, arr2.length <= 1000
 * 0 <= arr1[i], arr2[i] <= 1000
 * Each arr2[i] is distinct.
 * Each arr2[i] is in arr1.
 *
 *
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	count := [1001]int{}

	for _, a := range arr1 {
		count[a]++
	}

	fmt.Println(count)

	res := make([]int, 0, len(arr1))

	for _, b := range arr2 {
		for count[b] > 0 {
			res = append(res, b)
			count[b]--
		}
	}

	fmt.Println(res)

	for i := 0; i < 1001; i++ {
		for count[i] > 0 {
			res = append(res, i)
			count[i]--
		}
	}

	return res
}

// @lc code=end

