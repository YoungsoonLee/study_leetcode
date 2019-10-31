import "fmt"

/*
 * @lc app=leetcode id=876 lang=golang
 *
 * [876] Middle of the Linked List
 *
 * https://leetcode.com/problems/middle-of-the-linked-list/description/
 *
 * algorithms
 * Easy (65.07%)
 * Likes:    673
 * Dislikes: 46
 * Total Accepted:    90.3K
 * Total Submissions: 138.3K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given a non-empty, singly linked list with head node head, return a middle
 * node of linked list.
 *
 * If there are two middle nodes, return the second middle node.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: Node 3 from this list (Serialization: [3,4,5])
 * The returned node has value 3.  (The judge's serialization of this node is
 * [3,4,5]).
 * Note that we returned a ListNode object ans, such that:
 * ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next
 * = NULL.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4,5,6]
 * Output: Node 4 from this list (Serialization: [4,5,6])
 * Since the list has two middle nodes with values 3 and 4, we return the
 * second one.
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of nodes in the given list will be between 1 and 100.
 *
 *
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
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func my_middleNode(head *ListNode) *ListNode {
	temp := head
	size := 1
	for temp.Next != nil {
		size++
		temp = temp.Next
	}

	fmt.Println(size)

	temp = head
	for i := 0; i < (size / 2); i++ {
		temp = temp.Next
	}

	return temp
}

// @lc code=end

