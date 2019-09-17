/*
 * @lc app=leetcode id=628 lang=golang
 *
 * [628] Maximum Product of Three Numbers
 *
 * https://leetcode.com/problems/maximum-product-of-three-numbers/description/
 *
 * algorithms
 * Easy (46.55%)
 * Likes:    769
 * Dislikes: 285
 * Total Accepted:    81.5K
 * Total Submissions: 174.9K
 * Testcase Example:  '[1,2,3]'
 *
 * Given an integer array, find three numbers whose product is maximum and
 * output the maximum product.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3]
 * Output: 6
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [1,2,3,4]
 * Output: 24
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The length of the given array will be in range [3,10^4] and all elements are
 * in the range [-1000, 1000].
 * Multiplication of any three numbers in the input won't exceed the range of
 * 32-bit signed integer.
 * 
 * 
 * 
 * 
 */
func maximumProduct(nums []int) int {
	max, max1, max2 := -1001, -1001, -1001

	min1, min2 := 1001, 1001

	for _, n := range nums {
		switch {
		case n > max:
			max2, max1, max = max1, max, n
		case n > max1:
			max2, max1 = max1, n
		case n > max2:
			max2 = n
		}

		switch {
		case n < min1:
			min2, min1 = min1, n
		case n < min2:
			min2 = 2
		}
	}

	return bigger(max1*max2, min1*min2) * max

}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

