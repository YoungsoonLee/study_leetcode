/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 *
 * https://leetcode.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (30.41%)
 * Likes:    1905
 * Dislikes: 593
 * Total Accepted:    350.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '10'
 *
 * Count the number of prime numbers less than a non-negative number, n.
 * 
 * Example:
 * 
 * 
 * Input: 10
 * Output: 4
 * Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 * 
 * 
 */

// @lc code=start
func countPrimes(n int) int {
    if n < 3 {
		return 0
	}

	isComposite := make([]bool, n)
	// n보다 작은 수의 절반은 짝수, 2를 제외하고 prime이 아니다
	count := n/2

	for i:=3; i*i < n; i+=2 {
		if isComposite[i] {
			continue
		}

		for j := i*i; j<n; j+=2*i {
			if !isComposite[j] {
				count--
				isComposite[j] = true
			}
		}
	}

	return count
}
// @lc code=end

