/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 *
 * https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (52.25%)
 * Likes:    1455
 * Dislikes: 570
 * Total Accepted:    376.6K
 * Total Submissions: 709K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers that is already sorted in ascending order, find
 * two numbers such that they add up to a specific target number.
 * 
 * The function twoSum should return indices of the two numbers such that they
 * add up to the target, where index1 must be less than index2.
 * 
 * Note:
 * 
 * 
 * Your returned answers (both index1 and index2) are not zero-based.
 * You may assume that each input would have exactly one solution and you may
 * not use the same element twice.
 * 
 * 
 * Example:
 * 
 * 
 * Input: numbers = [2,7,11,15], target = 9
 * Output: [1,2]
 * Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 * 
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))

	for i,n := range numbers{
		if m[target-n] != 0 {
			return []int{m[target-n], i + 1}
		}
		m[n] = i+1	
	}
	return nil
}
// @lc code=end

