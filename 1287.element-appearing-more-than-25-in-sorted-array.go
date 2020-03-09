import "sort"

/*
 * @lc app=leetcode id=1287 lang=golang
 *
 * [1287] Element Appearing More Than 25% In Sorted Array
 *
 * https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/description/
 *
 * algorithms
 * Easy (60.42%)
 * Likes:    126
 * Dislikes: 8
 * Total Accepted:    11.4K
 * Total Submissions: 19.1K
 * Testcase Example:  '[1,2,2,6,6,6,6,7,10]'
 *
 * Given an integer array sorted in non-decreasing order, there is exactly one
 * integer in the array that occurs more than 25% of the time.
 *
 * Return that integer.
 *
 *
 * Example 1:
 * Input: arr = [1,2,2,6,6,6,6,7,10]
 * Output: 6
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 10^4
 * 0 <= arr[i] <= 10^5
 *
 */

// @lc code=start
func findSpecialInteger(arr []int) int {
	index, leng := 0, len(arr)
	t := leng >> 2 // 나누기 2
	for index < leng {
		curVal := arr[index]

		nextIndex := sort.Search(leng, func(i int) bool { return arr[i] > curVal })

		if nextIndex-index > t {
			return curVal
		}
		index = nextIndex
	}
	return -1
}

// @lc code=end

