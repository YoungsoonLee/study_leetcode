import "fmt"

/*
 * @lc app=leetcode id=859 lang=golang
 *
 * [859] Buddy Strings
 *
 * https://leetcode.com/problems/buddy-strings/description/
 *
 * algorithms
 * Easy (27.85%)
 * Likes:    370
 * Dislikes: 222
 * Total Accepted:    31.9K
 * Total Submissions: 114.8K
 * Testcase Example:  '"ab"\n"ba"'
 *
 * Given two strings A and B of lowercase letters, return true if and only if
 * we can swap two letters in A so that the result equals B.
 *
 *
 *
 * Example 1:
 *
 *
 *
 * Input: A = "ab", B = "ba"
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = "ab", B = "ab"
 * Output: false
 *
 *
 *
 * Example 3:
 *
 *
 * Input: A = "aa", B = "aa"
 * Output: true
 *
 *
 *
 * Example 4:
 *
 *
 * Input: A = "aaaaaaabc", B = "aaaaaaacb"
 * Output: true
 *
 *
 *
 * Example 5:
 *
 *
 * Input: A = "", B = "aa"
 * Output: false
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= A.length <= 20000
 * 0 <= B.length <= 20000
 * A and B consist only of lowercase letters.
 *
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if A == B {
		return hasDouble(A)
	}

	size := len(A)
	i := 0
	countDown := 2
	ca, cb := byte(0), byte(0)
	for countDown > 0 && i < size {
		if A[i] != B[i] {
			ca += A[i]
			cb += B[i]
			countDown--
		}
		i++
	}

	return ca == cv && A[i:] == B[i:]

}

func hasDouble(s string) bool {
	seen := [26]bool{}
	fmt.Println("seen: ", seen)

	for i := range s {
		b := s[i] - 'a'
		if seen[b] {
			fmt.Println("seen true: ", seen)
			return true
		}
	}
	fmt.Println("seen false: ", seen)
	return false
}

// @lc code=end

