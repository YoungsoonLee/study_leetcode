/*
 * @lc app=leetcode id=784 lang=golang
 *
 * [784] Letter Case Permutation
 *
 * https://leetcode.com/problems/letter-case-permutation/description/
 *
 * algorithms
 * Easy (58.85%)
 * Likes:    787
 * Dislikes: 92
 * Total Accepted:    59.9K
 * Total Submissions: 100.6K
 * Testcase Example:  '"a1b2"'
 *
 * Given a string S, we can transform every letter individually to be lowercase
 * or uppercase to create another string.  Return a list of all possible
 * strings we could create.
 *
 *
 * Examples:
 * Input: S = "a1b2"
 * Output: ["a1b2", "a1B2", "A1b2", "A1B2"]
 *
 * Input: S = "3z4"
 * Output: ["3z4", "3Z4"]
 *
 * Input: S = "12345"
 * Output: ["12345"]
 *
 *
 * Note:
 *
 *
 * S will be a string with length between 1 and 12.
 * S will consist only of letters or digits.
 *
 *
 */

// @lc code=start

func letterCasePermutation(S string) []string {
	size := len(s)
	if size == 0 {
		return []string{""}
	}

	// s의 마지막 문자를 접미사로 추출
	lastByte := s[size-1]
	postfixs := make([]string, 1, 2)
	postfixs[0] = string(lastByte)

	// 마지막 문자가 문자 인 경우
	// 접미사, 다른 형태의 글쓰기 추가
	if b, ok := check(lastByte); ok {
		postfixs = append(postfixs, string(b))
	}

	prefixs := letterCasePermutation(s[:size-1])

	// res는 접두사와 접미사의 곱입니다.
	res := make([]string, 0, len(prefixs)*len(prefixs))

	for _, pre := range prefixs {
		for _, post := range postfixs {
			res = append(res, pre+post)
		}
	}

	return res
}

func check(b byte) (byte, bool) {
	if 'a' <= b && b <= 'z' {
		return b + 'A' - 'a', true
	}
	if 'A' <= b && b <= 'Z' {
		return b + 'a' - 'A', true
	}
	return 0, false
}

// @lc code=end

