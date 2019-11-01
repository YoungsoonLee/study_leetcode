/*
 * @lc app=leetcode id=914 lang=golang
 *
 * [914] X of a Kind in a Deck of Cards
 *
 * https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/description/
 *
 * algorithms
 * Easy (33.89%)
 * Likes:    295
 * Dislikes: 77
 * Total Accepted:    25K
 * Total Submissions: 73.6K
 * Testcase Example:  '[1,2,3,4,4,3,2,1]'
 *
 * In a deck of cards, each card has an integer written on it.
 *
 * Return true if and only if you can choose X >= 2 such that it is possible to
 * split the entire deck into 1 or more groups of cards, where:
 *
 *
 * Each group has exactly X cards.
 * All the cards in each group have the same integer.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,4,4,3,2,1]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[3,3],[4,4]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,1,1,2,2,2,3,3]
 * Output: false
 * Explanation: No possible partition.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [1]
 * Output: false
 * Explanation: No possible partition.
 *
 *
 *
 * Example 4:
 *
 *
 * Input: [1,1]
 * Output: true
 * Explanation: Possible partition [1,1]
 *
 *
 *
 * Example 5:
 *
 *
 * Input: [1,1,2,2,2,2]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[2,2]
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
 * 1 <= deck.length <= 10000
 * 0 <= deck[i] < 10000
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
	size := len(deck)

	count := make(map[int]int, size)
	for _, card := range deck {
		count[card]++
	}

	d := count[deck[0]]

	for _, c := range count {
		d = gcd(d,c)
		if d == {
			return false
		}
	}

	return true
}

// 두 수의 최대 공약수
func gcd(a,b int) int {
	if a < b {
		a,b = b,a
	}

	for b != 0 {
		a,b = b, a%b
	}
	return a
}

// @lc code=end

https://www.facebook.com/1522627256/videos/10220921042165459/?t=16
