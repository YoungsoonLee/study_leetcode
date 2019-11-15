import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=933 lang=golang
 *
 * [933] Number of Recent Calls
 *
 * https://leetcode.com/problems/number-of-recent-calls/description/
 *
 * algorithms
 * Easy (69.42%)
 * Likes:    176
 * Dislikes: 975
 * Total Accepted:    32.2K
 * Total Submissions: 46.1K
 * Testcase Example:  '["RecentCounter","ping","ping","ping","ping"]\n[[],[1],[100],[3001],[3002]]'
 *
 * Write a class RecentCounter to count recent requests.
 *
 * It has only one method: ping(int t), where t represents some time in
 * milliseconds.
 *
 * Return the number of pings that have been made from 3000 milliseconds ago
 * until now.
 *
 * Any ping with time in [t - 3000, t] will count, including the current ping.
 *
 * It is guaranteed that every call to ping uses a strictly larger value of t
 * than before.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: inputs = ["RecentCounter","ping","ping","ping","ping"], inputs =
 * [[],[1],[100],[3001],[3002]]
 * Output: [null,1,2,3,3]
 *
 *
 *
 * Note:
 *
 *
 * Each test case will have at most 10000 calls to ping.
 * Each test case will call ping with strictly increasing values of t.
 * Each call to ping will have 1 <= t <= 10^9.
 *
 *
 *
 *
 *
 */

// @lc code=start
type RecentCounter struct {
	times []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		times: make([]int, 0, 10000),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.times = append(this.times, t)
	//fmt.Println(this.times)
	fmt.Println(sort.SearchInts(this.times, t-3000))
	return len(this.times) - sort.SearchInts(this.times, t-3000)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

