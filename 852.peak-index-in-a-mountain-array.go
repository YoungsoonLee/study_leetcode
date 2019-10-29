/*
 * @lc app=leetcode id=852 lang=golang
 *
 * [852] Peak Index in a Mountain Array
 *
 * https://leetcode.com/problems/peak-index-in-a-mountain-array/description/
 *
 * algorithms
 * Easy (70.08%)
 * Likes:    375
 * Dislikes: 821
 * Total Accepted:    104.7K
 * Total Submissions: 149K
 * Testcase Example:  '[0,1,0]'
 *
 * Let's call an array A a mountain if the following properties hold:
 *
 *
 * A.length >= 3
 * There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] <
 * A[i] > A[i+1] > ... > A[A.length - 1]
 *
 *
 * Given an array that is definitely a mountain, return any i such that A[0] <
 * A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].
 *
 * Example 1:
 *
 *
 * Input: [0,1,0]
 * Output: 1
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [0,2,1,0]
 * Output: 1
 *
 *
 * Note:
 *
 *
 * 3 <= A.length <= 10000
 * 0 <= A[i] <= 10^6
 * A is a mountain, as defined above.
 *
 *
 */

// @lc code=start
func peakIndexInMountainArray(A []int) int {
	/*
		if len(A) < 3 {
			// not mountain
			return 0
		}

		peakIndex := 0
		for i := 1; i < len(A); i++ {
			if A[i-1] < A[i] {
				peakIndex = i
			}
		}

		return peakIndex
	*/
	l, r := 0, len(A)-1

	for {
		m := (l + r) / 2
		switch {
		case a[m] < a[m+1]:
			l = m
		case a[m-1] > a[m]:
			r = m
		default:
			return m
		}
	}
}

// @lc code=end

