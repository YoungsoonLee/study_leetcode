import "fmt"

/*
 * @lc app=leetcode id=989 lang=golang
 *
 * [989] Add to Array-Form of Integer
 *
 * https://leetcode.com/problems/add-to-array-form-of-integer/description/
 *
 * algorithms
 * Easy (43.92%)
 * Likes:    217
 * Dislikes: 43
 * Total Accepted:    28.5K
 * Total Submissions: 64.6K
 * Testcase Example:  '[1,2,0,0]\n34'
 *
 * For a non-negative integer X, the array-form of X is an array of its digits
 * in left to right order.  For example, if X = 1231, then the array form is
 * [1,2,3,1].
 *
 * Given the array-form A of a non-negative integer X, return the array-form of
 * the integer X+K.
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
 * Input: A = [1,2,0,0], K = 34
 * Output: [1,2,3,4]
 * Explanation: 1200 + 34 = 1234
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = [2,7,4], K = 181
 * Output: [4,5,5]
 * Explanation: 274 + 181 = 455
 *
 *
 *
 * Example 3:
 *
 *
 * Input: A = [2,1,5], K = 806
 * Output: [1,0,2,1]
 * Explanation: 215 + 806 = 1021
 *
 *
 *
 * Example 4:
 *
 *
 * Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
 * Output: [1,0,0,0,0,0,0,0,0,0,0]
 * Explanation: 9999999999 + 1 = 10000000000
 *
 *
 *
 *
 * Note：
 *
 *
 * 1 <= A.length <= 10000
 * 0 <= A[i] <= 9
 * 0 <= K <= 10000
 * If A.length > 1, then A[0] != 0
 *
 *
 *
 *
 *
 */

// @lc code=start

func addToArrayForm(A []int, K int) []int {
	j := 1
	sum := 0
	for i := len(A) - 1; i >= 0; i-- {
		sum += A[i] * j
		j *= 10

	}

	sum += K
	fmt.Println(sum)
	fmt.Println(numToArray(sum))
	return []int{}
}

func numToArray(a int) []int {
	res := make([]int, 0, 8)
	for a > 0 {
		res = append(res, a%10) // 나머지
		a /= 10                 // 몫업데이트
	}

	reverse(res)
	return res
}

func reverse(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

// @lc code=end

