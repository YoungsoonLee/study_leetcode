/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (50.71%)
 * Likes:    3542
 * Dislikes: 524
 * Total Accepted:    886.2K
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// using channel

	var ret, last *ListNode
	min := make(chan *ListNode) // unbuffered channel

	go func() {
		defer close(min)

		for {
			if l1 == nil {
				min <- l2
				return
			}

			if l2 == nil {
				min <- l1
				return
			}

			if l1.Val < l2.Val {
				min <- l1
				l1 = l1.Next
			} else {
				min <- l2
				l2 = l2.Next
			}

		}
	}()

	// consume the channel
	for v := range min {
		if last != nil {
			last.Next = v
		}
		if ret == nil {
			ret = v
		}
		last = v
	}

	return ret
}

// @lc code=end

