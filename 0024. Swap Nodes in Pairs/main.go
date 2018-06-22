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
	if head == nil || head.Next == nil {
		return head
	}

	// 머리글을 가리키는 newHead.Next 노드
	newHead := head.Next

	// head.Next 노드를 newHead로 변환합니다.
	head.Next = swapPairs(newHead.Next)
	// 헤드 노드를 가리킨 다음 newHead를 보자.
	newHead.Next = head

	return newHead
}

func swapPairs_my(head *ListNode) *ListNode {
	fast := head.Next

	for fast != nil {
		head.Val, fast.Val = fast.Val, head.Val
		head = fast.Next
		fast = fast.Next.Next
	}

	return head

}

func main() {}
