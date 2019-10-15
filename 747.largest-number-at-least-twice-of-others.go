/*
 * @lc app=leetcode id=747 lang=golang
 *
 * [747] Largest Number At Least Twice of Others
 *
 * https://leetcode.com/problems/largest-number-at-least-twice-of-others/description/
 *
 * algorithms
 * Easy (40.80%)
 * Likes:    227
 * Dislikes: 455
 * Total Accepted:    62.6K
 * Total Submissions: 153K
 * Testcase Example:  '[0,0,0,1]'
 *
 * In a given integer array nums, there is always exactly one largest element.
 *
 * Find whether the largest element in the array is at least twice as much as
 * every other number in the array.
 *
 * If it is, return the index of the largest element, otherwise return -1.
 *
 * Example 1:
 *
 *
 * Input: nums = [3, 6, 1, 0]
 * Output: 1
 * Explanation: 6 is the largest integer, and for every other number in the
 * array x,
 * 6 is more than twice as big as x.  The index of value 6 is 1, so we return
 * 1.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1, 2, 3, 4]
 * Output: -1
 * Explanation: 4 isn't at least as big as twice the value of 3, so we return
 * -1.
 *
 *
 *
 *
 * Note:
 *
 *
 * nums will have a length in the range [1, 50].
 * Every nums[i] will be an integer in the range [0, 99].
 *
 *
 *
 *
 */

// @lc code=start
func dominantIndex(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	i1, i2 := 0, 1 // index
	if nums[i1] < nums[i2] {
		i1, i2 = i2, i1
	}

	for i := 2; i < n; i++ {
		if nums[i1] < nums[i] {
			i2, i1 = i1, i
		} else if nums[i2] < nums[i] {
			i2 = i
		}
	}

	if nums[i2] == 0 || nums[i1]/a[i2] >= 2 {
		return i1
	}

	return -1

}

// @lc code=end

