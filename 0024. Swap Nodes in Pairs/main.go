package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	fast := head.Next

	for fast != nil {
		head.Val, fast.Val = fast.Val, head.Val
		head = fast.Next
		fast = fast.Next.Next
	}

	return head

}

func main() {}
