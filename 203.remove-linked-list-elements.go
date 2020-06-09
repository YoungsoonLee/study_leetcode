/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (36.91%)
 * Likes:    1454
 * Dislikes: 85
 * Total Accepted:    314.2K
 * Total Submissions: 837.3K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 * 
 * Example:
 * 
 * 
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
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
func removeElements(head *ListNode, val int) *ListNode {

	headPre := ListNode{Next: head}

	temp := &headPre

	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		}else{
			temp = temp.Next
		}
	}

	return headPre.Next


}
// @lc code=end

