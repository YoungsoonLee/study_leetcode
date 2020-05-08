/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (55.40%)
 * Likes:    2841
 * Dislikes: 213
 * Total Accepted:    588.4K
 * Total Submissions: 1M
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 * 
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,2,3]
 * Output: 3
 * 
 * Example 2:
 * 
 * 
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 * 
 * 
 */

// @lc code=start
func majorityElement(nums []int) int {
	x, t := nums[0], 1

	for i:=1; i<len(nums); i++ {
		fmt.Println(x, t)
		
		switch {
		case x == nums[i]:
			t++
		case t>0:
			t--
		default:
			x = nums[i]
			t = 1
		}
	}

	return x
}
// @lc code=end

