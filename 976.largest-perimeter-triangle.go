import "container/heap"

/*
 * @lc app=leetcode id=976 lang=golang
 *
 * [976] Largest Perimeter Triangle
 *
 * https://leetcode.com/problems/largest-perimeter-triangle/description/
 *
 * algorithms
 * Easy (56.91%)
 * Likes:    241
 * Dislikes: 27
 * Total Accepted:    23.8K
 * Total Submissions: 41.7K
 * Testcase Example:  '[2,1,2]'
 *
 * Given an array A of positive lengths, return the largest perimeter of a
 * triangle with non-zero area, formed from 3 of these lengths.
 *
 * If it is impossible to form any triangle of non-zero area, return 0.
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
 * Input: [2,1,2]
 * Output: 5
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,1]
 * Output: 0
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [3,2,3,4]
 * Output: 10
 *
 *
 *
 * Example 4:
 *
 *
 * Input: [3,6,2,3]
 * Output: 8
 *
 *
 *
 *
 * Note:
 *
 *
 * 3 <= A.length <= 10000
 * 1 <= A[i] <= 10^6
 *
 *
 *
 *
 *
 */

// @lc code=start

func largestPerimeter(A []int) int {
	size := len(A)
	h := intHeap(A)

	heap.Init(&h)

	a := heap.Pop(&h).(int)
	b := heap.Pop(&h).(int)
	for i := size - 3; i >= 0; i-- {
		c := heap.Pop(&h).(int)
		if a < b+c {
			return a + b + c
		}
		a, b = b, c
	}
	return 0
}

// intHeap 实现了 heap 的接口
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j] // NOTICE: Max is at the top
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

// @lc code=end

