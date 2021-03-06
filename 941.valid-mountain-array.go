/*
 * @lc app=leetcode id=941 lang=golang
 *
 * [941] Valid Mountain Array
 *
 * https://leetcode.com/problems/valid-mountain-array/description/
 *
 * algorithms
 * Easy (35.22%)
 * Likes:    221
 * Dislikes: 54
 * Total Accepted:    30.3K
 * Total Submissions: 86.2K
 * Testcase Example:  '[2,1]'
 *
 * Given an array A of integers, return true if and only if it is a valid
 * mountain array.
 *
 * Recall that A is a mountain array if and only if:
 *
 *
 * A.length >= 3
 * There exists some i with 0 < i < A.length - 1 such that:
 *
 * A[0] < A[1] < ... A[i-1] < A[i]
 * A[i] > A[i+1] > ... > A[A.length - 1]
 *
 *
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [2,1]
 * Output: false
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [3,5,5]
 * Output: false
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [0,3,2,1]
 * Output: true
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= A.length <= 10000
 * 0 <= A[i] <= 10000
 *
 *
 *
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func validMountainArray(A []int) bool {
	size := len(A)

	if size < 3 {
		return false
	}

	i := 1
	for i < size && A[i-1] < A[i] {
		i++
	}

	top := i - 1 // top position

	for i < size && A[i-1] > A[i] {
		i++
	}

	return 0 < top && top < size-1 && i == size
}

// @lc code=end

