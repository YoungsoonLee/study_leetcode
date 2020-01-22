/*
 * @lc app=leetcode id=1005 lang=golang
 *
 * [1005] Maximize Sum Of Array After K Negations
 *
 * https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/
 *
 * algorithms
 * Easy (50.57%)
 * Likes:    220
 * Dislikes: 27
 * Total Accepted:    19.9K
 * Total Submissions: 39.3K
 * Testcase Example:  '[4,2,3]\n1'
 *
 * Given an array A of integers, we must modify the array in the following way:
 * we choose an i and replace A[i] with -A[i], and we repeat this process K
 * times in total.  (We may choose the same index i multiple times.)
 * 
 * Return the largest possible sum of the array after modifying it in this
 * way.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: A = [4,2,3], K = 1
 * Output: 5
 * Explanation: Choose indices (1,) and A becomes [4,-2,3].
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: A = [3,-1,0,2], K = 3
 * Output: 6
 * Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: A = [2,-3,-1,5,-4], K = 2
 * Output: 13
 * Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
 * 
 * 
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= A.length <= 10000
 * 1 <= K <= 10000
 * -100 <= A[i] <= 100
 * 
 * 
 */

// @lc code=start
func largestSumAfterKNegations(A []int, K int) int {
	size := len(A)
	
	i,j := 0, size-1
	minAbs := abs(A[0])

	for i<=j {
		minAbs = min(minAbs, abs(A[i]))
		if A[i] >= 0 { // move non-negative to right side
			A[i], A[j] = A[j], A[i]
			j--
			continue
		}
		i++
	}

	negatives := A[:j+1]
	negSize := j+1

	sum := 0

	if K >= negSize { // all negative could convert to positive
		if (K-negSize)&1 == 1{ // someone need keep as negative
			sum -= minAbs * 2 // choose minAbs keep as negative
		}
	}else{
		sort.Ints(negatives)
		// min K negative could convert to positive
		// sort make min K negative in negatives[:K]
	}

	index := min(K, negSize)
	for i:=0; i<index; i++ {
		sum -= A[i]
	}

	for i:=index; i<size; i++ {
		sum += A[i]
	}

	return sum

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

