/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (41.76%)
 * Likes:    1537
 * Dislikes: 258
 * Total Accepted:    413.4K
 * Total Submissions: 961.8K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 * 
 * The input strings are both non-empty and contains only characters 1 or 0.
 * 
 * Example 1:
 * 
 * 
 * Input: a = "11", b = "1"
 * Output: "100"
 * 
 * Example 2:
 * 
 * 
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 * 
 * 
 * Constraints:
 * 
 * 
 * Each string consists only of '0' or '1' characters.
 * 1 <= a.length, b.length <= 10^4
 * Each string is either "0" or doesn't contain any leading zero.
 * 
 * 
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	l := len(a)

	isA := trans(a, l)
	isB := trans(b, l)

	return makeString(isA, isB)
}

func trans(s string, l int) []int {
	res := make([]int, l)
	ls := len(s)
	for i, b := range s {
		res[l-ls+i] = int(b-'0')
	}
	return res
}

func add(a, b []int) []int {
	l := len(a)+1
	res := make([]int, l)
	for i:=l-1; i>=1; i-- {
		temp := res[i] +b[i-1]
		res[i] = temp %2
		res[i-1] = temp /2
	}

	i := 0
	for i<l-1 && res[i] == 0{
		i++
	}

	return res[i:]
}

func makeString(nums []int) string {
	bytes := make([]byte, len(nums))
	for i:= range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}
	return string(bytes)
}
// @lc code=end

