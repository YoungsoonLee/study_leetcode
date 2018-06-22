package main

import "fmt"

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

	fmt.Println("head: ", head)
	fmt.Println("new: ", newHead)

	return newHead
}

func main() {
	l4 := &ListNode{Val: 4, Next: nil}
	l3 := &ListNode{Val: 3, Next: l4}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 1, Next: l2}
	r := swapPairs(l1)
	fmt.Println(r)

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
