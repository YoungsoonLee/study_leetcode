/*
 * @lc app=leetcode id=942 lang=golang
 *
 * [942] DI String Match
 *
 * https://leetcode.com/problems/di-string-match/description/
 *
 * algorithms
 * Easy (69.87%)
 * Likes:    578
 * Dislikes: 215
 * Total Accepted:    46.7K
 * Total Submissions: 66.6K
 * Testcase Example:  '"IDID"'
 *
 * Given a string S that only contains "I" (increase) or "D" (decrease), let N
 * = S.length.
 *
 * Return any permutation A of [0, 1, ..., N] such that for all i = 0, ...,
 * N-1:
 *
 *
 * If S[i] == "I", then A[i] < A[i+1]
 * If S[i] == "D", then A[i] > A[i+1]
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "IDID"
 * Output: [0,4,1,3,2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "III"
 * Output: [0,1,2,3]
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "DDI"
 * Output: [3,2,0,1]
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= S.length <= 10000
 * S only contains characters "I" or "D".
 *
 */

// @lc code=start
func diStringMatch(S string) []int {
	size := len(S)
	l, r := 0, size

	a := make([]int, size+1)

	for i, b := range S {
		if b == 'I' {
			a[i] = l
			l++
		} else {
			a[i] = r
			r--
		}
		i++
	}

	a[size] = l

	return a

	/* my solution
	size := len(S)
	a := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		a[i] = i
	}

	fmt.Println(a)

	result := make([]int, 0, size+1)
	for _, c := range S {
		switch c {
		case 'I':
			result = append(result, a[0])
			a = a[1:]
		case 'D':
			result = append(result, a[len(a)-1])
			a = a[:len(a)-1]
		}

	}

	result = append(result, a...)
	fmt.Println(result)

	return result
	*/
}

// @lc code=end

