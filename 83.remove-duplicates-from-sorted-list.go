/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (44.01%)
 * Likes:    1277
 * Dislikes: 101
 * Total Accepted:    432.1K
 * Total Submissions: 967.7K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 * 
 * Example 1:
 * 
 * 
 * Input: 1->1->2
 * Output: 1->2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 1->1->2->3->3
 * Output: 1->2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	next := head.Next

	for next.Next != nil {
		if cur.Val == next.Val {
			cur.Next = next.Next
		}
		cur = next
		next = next.Next
	}

	// last
	if cur.Val == next.Val {
		cur = nil
	}

	return head
}
// @lc code=end

