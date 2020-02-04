/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 *
 * https://leetcode.com/problems/last-stone-weight/description/
 *
 * algorithms
 * Easy (62.52%)
 * Likes:    344
 * Dislikes: 18
 * Total Accepted:    38.1K
 * Total Submissions: 60.8K
 * Testcase Example:  '[2,7,4,1,8,1]'
 *
 * We have a collection of rocks, each rock has a positive integer weight.
 * 
 * Each turn, we choose the two heaviest rocks and smash them together.
 * Suppose the stones have weights x and y with x <= y.  The result of this
 * smash is:
 * 
 * 
 * If x == y, both stones are totally destroyed;
 * If x != y, the stone of weight x is totally destroyed, and the stone of
 * weight y has new weight y-x.
 * 
 * 
 * At the end, there is at most 1 stone left.  Return the weight of this stone
 * (or 0 if there are no stones left.)
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: [2,7,4,1,8,1]
 * Output: 1
 * Explanation: 
 * We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
 * we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
 * we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
 * we combine 1 and 1 to get 0 so the array converts to [1] then that's the
 * value of last stone.
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= stones.length <= 30
 * 1 <= stones[i] <= 1000
 * 
 */

// @lc code=start
import "container/heap"
func lastStoneWeight(stones []int) int {
	
	pq := PQ(stones)
	heap.Init(&pq) // !!! 힙을 이용. 매번 정렬을 해야 하는 것을 막는다. !!!!!

	for len(pq) >= 2 {
		a, b := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
		if a == b {
			continue
		}
		heap.Push(&pq, a-b)
	}

	if len(pq) == 0 {
		return 0
	}

	return pq[0]
	
}

// PQ implements heap.Interface and holds entries.
type PQ []int

func (pq PQ) Len() int{return len(pq)}

func (pq PQ) Less(i,j int) bool {
	return pq[i] > pq[j]
}

func (pq PQ) Swap(i, j int) { 
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	temp := x.(int)
	*pq = append(*pq, temp)
}

func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0:len(*pq)-1]
	return temp
}

// @lc code=end

