/*
 * @lc app=leetcode id=1207 lang=golang
 *
 * [1207] Unique Number of Occurrences
 *
 * https://leetcode.com/problems/unique-number-of-occurrences/description/
 *
 * algorithms
 * Easy (71.73%)
 * Likes:    234
 * Dislikes: 15
 * Total Accepted:    37.3K
 * Total Submissions: 52.3K
 * Testcase Example:  '[1,2,2,1,1,3]'
 *
 * Given an array of integers arr, write a function that returns true if and
 * only if the number of occurrences of each value in the array is unique.
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [1,2,2,1,1,3]
 * Output: true
 * Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two
 * values have the same number of occurrences.
 *
 * Example 2:
 *
 *
 * Input: arr = [1,2]
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 1000
 * -1000 <= arr[i] <= 1000
 *
 *
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int, len(arr))
	for _, a := range arr {
		m[a]++
	}

	hasSeen := make(map[int]bool, len(m))
	for _, c := range m {
		if hasSeen[c] {
			return false
		}
		hasSeen[c] = true
	}

	return true

	/*
		res := make([]int, 0)
		for _, c := range m {
			res = append(res, c)
		}

		sort.Ints(res)

		for i := 0; i < len(res)-1; i++ {
			if res[i] == res[i+1] {
				return false
			}
		}

		return true
	*/
}

// @lc code=end

