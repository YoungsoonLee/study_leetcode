/*
 * @lc app=leetcode id=788 lang=golang
 *
 * [788] Rotated Digits
 *
 * https://leetcode.com/problems/rotated-digits/description/
 *
 * algorithms
 * Easy (55.29%)
 * Likes:    240
 * Dislikes: 815
 * Total Accepted:    35.3K
 * Total Submissions: 63.6K
 * Testcase Example:  '10'
 *
 * X is a good number if after rotating each digit individually by 180 degrees,
 * we get a valid number that is different from X.  Each digit must be rotated
 * - we cannot choose to leave it alone.
 *
 * A number is valid if each digit remains a digit after rotation. 0, 1, and 8
 * rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each
 * other, and the rest of the numbers do not rotate to any other number and
 * become invalid.
 *
 * Now given a positive number N, how many numbers X from 1 to N are good?
 *
 *
 * Example:
 * Input: 10
 * Output: 4
 * Explanation:
 * There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
 * Note that 1 and 10 are not good numbers, since they remain unchanged after
 * rotating.
 *
 *
 * Note:
 *
 *
 * N  will be in range [1, 10000].
 *
 *
 */

// @lc code=start
func rotatedDigits(N int) int {
	count := 0
	//brute force
	for i := 2; i <= N; i++ {
		if isValid(i) {
			count++
		}
	}

	return count
}

func isValid(n int) bool {
	var hasFoundValid bool
	for n > 0 {
		switch n % 10 {
		case 2, 5, 6, 9:
			hasFoundValid = true
		case 3, 4, 7:
			return false
		}

		n /= 10
	}

	return hasFoundValid
}

// @lc code=end

