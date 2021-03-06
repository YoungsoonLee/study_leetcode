/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (44.94%)
 * Likes:    13848
 * Dislikes: 509
 * Total Accepted:    2.6M
 * Total Submissions: 5.9M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, b := range nums {
		if j, ok := m[target-b]; ok {
			return []int{j, i}
		}	
		m[b] = i
	}
	return nil
}

// @lc code=end

