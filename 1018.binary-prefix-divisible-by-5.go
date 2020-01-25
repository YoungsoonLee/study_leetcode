/*
 * @lc app=leetcode id=1018 lang=golang
 *
 * [1018] Binary Prefix Divisible By 5
 *
 * https://leetcode.com/problems/binary-prefix-divisible-by-5/description/
 *
 * algorithms
 * Easy (47.05%)
 * Likes:    154
 * Dislikes: 75
 * Total Accepted:    17K
 * Total Submissions: 36.2K
 * Testcase Example:  '[0,1,1]'
 *
 * Given an array A of 0s and 1s, consider N_i: the i-th subarray from A[0] to
 * A[i] interpreted as a binary number (from most-significant-bit to
 * least-significant-bit.)
 *
 * Return a list of booleans answer, where answer[i] is true if and only if N_i
 * is divisible by 5.
 *
 * Example 1:
 *
 *
 * Input: [0,1,1]
 * Output: [true,false,false]
 * Explanation:
 * The input numbers in binary are 0, 01, 011; which are 0, 1, and 3 in
 * base-10.  Only the first number is divisible by 5, so answer[0] is true.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,1,1]
 * Output: [false,false,false]
 *
 *
 * Example 3:
 *
 *
 * Input: [0,1,1,1,1,1]
 * Output: [true,false,false,false,true,false]
 *
 *
 * Example 4:
 *
 *
 * Input: [1,1,1,0,1]
 * Output: [false,false,false,false,false]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 30000
 * A[i] is 0 or 1
 *
 *
 */

// @lc code=start
func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	r := 0

	for i, a := range A {
		r = (r*2 + a) % 5
		res[i] = r == 0
	}

	return res

}

// @lc code=end

