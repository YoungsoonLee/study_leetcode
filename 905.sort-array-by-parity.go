/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 *
 * https://leetcode.com/problems/sort-array-by-parity/description/
 *
 * algorithms
 * Easy (72.77%)
 * Likes:    594
 * Dislikes: 62
 * Total Accepted:    137.8K
 * Total Submissions: 188.8K
 * Testcase Example:  '[3,1,2,4]'
 *
 * Given an array A of non-negative integers, return an array consisting of all
 * the even elements of A, followed by all the odd elements of A.
 *
 * You may return any answer array that satisfies this condition.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [3,1,2,4]
 * Output: [2,4,3,1]
 * The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 5000
 * 0 <= A[i] <= 5000
 *
 *
 *
 */

// @lc code=start
func sortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1
	for {
		for i < j && A[i]%2 == 0 {
			i++
		}
		for i < j && A[j]%2 == 1 {
			j--
		}
		if i < j {
			A[i], A[j] = A[j], A[i]
		} else {
			break
		}
	}
	return A

	/*
		* 메모리를 너무 많이 쓰는 방식.
		size := len(A)
		rec := make([]int, 0, size)
		tmpOdd := make([]int, 0, size)

		for i := 0; i < size; i++ {
			if (A[i] % 2) != 0 {
				tmpOdd = append(tmpOdd, A[i])
			} else {
				rec = append(rec, A[i])
			}
		}

		rec = append(rec, tmpOdd...)
		return rec
	*/
}

// @lc code=end

