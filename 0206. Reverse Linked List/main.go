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

func (l *ListNode) print() {
	for l != nil {
		fmt.Printf("%d -> ", l.Val)
		l = l.Next
	}
}

// !!!
func (l *ListNode) add(val int) {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ListNode{Val: val, Next: nil} // !!!!!
}

func reverse_1(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head //!!
		head = temp
	}

	return prev
}

// !!!!!!!!!!
func reverseList(head *ListNode) *ListNode {
	// prev는 모든 되 돌린 노드의 머리 부분이다.
	var prev *ListNode
	for head != nil {
		head.print()
		fmt.Println()
		//// temp가 head를 가리 키도록합니다. 다음으로, head. 다음은 사라졌습니다.
		temp := head.Next
		// head는 복귀 된 노드의 새로운 헤드라고 불린다.\
		head.Next = prev // !!!
		head.print()
		fmt.Println()

		prev = head
		head = temp
	}

	return prev
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return p
}

func main() {
	a1 := &ListNode{Val: 1, Next: nil}
	//a1.print()
	a1.add(2)
	a1.add(3)
	a1.add(4)
	a1.add(5)
	//a1.print()
	k := reverseList(a1)
	k.print()
	//a1.print()
	//reverseList2(a1)
	//a1.print()

}
