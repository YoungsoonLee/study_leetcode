/*
 * @lc app=leetcode id=893 lang=golang
 *
 * [893] Groups of Special-Equivalent Strings
 *
 * https://leetcode.com/problems/groups-of-special-equivalent-strings/description/
 *
 * algorithms
 * Easy (63.51%)
 * Likes:    173
 * Dislikes: 872
 * Total Accepted:    19.2K
 * Total Submissions: 30.1K
 * Testcase Example:  '["abcd","cdab","cbad","xyzz","zzxy","zzyx"]'
 *
 * You are given an array A of strings.
 *
 * A move onto S consists of swapping any two even indexed characters of S, or
 * any two odd indexed characters of S.
 *
 * Two strings S and T are special-equivalent if after any number of moves onto
 * S, S == T.
 *
 * For example, S = "zzxy" and T = "xyzz" are special-equivalent because we may
 * make the moves "zzxy" -> "xzzy" -> "xyzz" that swap S[0] and S[2], then S[1]
 * and S[3].
 *
 * Now, a group of special-equivalent strings from A is a non-empty subset of A
 * such that:
 *
 *
 * Every pair of strings in the group are special equivalent, and;
 * The group is the largest size possible (ie., there isn't a string S not in
 * the group such that S is special equivalent to every string in the group)
 *
 *
 * Return the number of groups of special-equivalent strings from A.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
 * Output: 3
 * Explanation:
 * One group is ["abcd", "cdab", "cbad"], since they are all pairwise special
 * equivalent, and none of the other strings are all pairwise special
 * equivalent to these.
 *
 * The other two groups are ["xyzz", "zzxy"] and ["zzyx"].  Note that in
 * particular, "zzxy" is not special equivalent to "zzyx".
 *
 *
 *
 * Example 2:
 *
 *
 * Input: ["abc","acb","bac","bca","cab","cba"]
 * Output: 3
 *
 *
 *
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 1000
 * 1 <= A[i].length <= 20
 * All A[i] have the same length.
 * All A[i] consist of only lowercase letters.
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func numSpecialEquivGroups(A []string) int {
	strSize := len(A[0])
	groups := make(map[[26]int]bool, len(A))

	for _, a := range A {
		count := [26]int{}
		i := 0
		for i := 0; i+1 < strSize; i += 2 {
			// 패리티를 결정하기 위해 mod 연산을 피하려면 루프 확장을 사용하십시오.
			// 또한 중요한 경로의 길이를 줄였습니다.
			count[a[i]-'a']++
			// 조건부 조건으로 인해 strSize <= 20
			// 따라서 패리티의 글자를 동시에 셀 수 있습니다.
			count[a[i+1]-'a'] += 100
		}

		if i < strSize {
			count[a[i]-'a']++
		}
		groups[count] = true
	}
	return len(groups)
}

// @lc code=end

