/*
 * @lc app=leetcode id=1013 lang=golang
 *
 * [1013] Partition Array Into Three Parts With Equal Sum
 *
 * https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/description/
 *
 * algorithms
 * Easy (56.90%)
 * Likes:    262
 * Dislikes: 45
 * Total Accepted:    24.4K
 * Total Submissions: 42.9K
 * Testcase Example:  '[0,2,1,-6,6,-7,9,1,2,0,1]'
 *
 * Given an array A of integers, return true if and only if we can partition
 * the array into three non-empty parts with equal sums.
 * 
 * Formally, we can partition the array if we can find indexes i+1 < j with
 * (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1]
 * + ... + A[A.length - 1])
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: [0,2,1,-6,6,-7,9,1,2,0,1]
 * Output: true
 * Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [0,2,1,-6,6,7,9,-1,2,0,1]
 * Output: false
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: [3,3,6,5,-2,2,5,1,-9,4]
 * Output: true
 * Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
 * 
 * 
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 3 <= A.length <= 50000
 * -10000 <= A[i] <= 10000
 * 
 */

// @lc code=start
func canThreePartsEqualSum(A []int) bool {
	n := len(A)
	sums := make([]int, n)
	s := 0

	for i:=0; i<n; i++ {
		s += A[i]
		sums[i] = s
	}

	//fmt.Println(sums)
	//fmt.Println(s)
	if s%3 != 0 {
		return false
	}

	s /= 3
	i := 0
	for i<n && sums[i] != s {
		i++
	}

	s *= 2
	j := n-1
	for 0<=j && sums[j] != s {
		j--
	}

	return i < j

}
// @lc code=end

