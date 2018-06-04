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

func newList(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func (l *ListNode) add(val int) {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ListNode{Val: val, Next: nil}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	for head != nil {
		if head.Next != nil {
			if head.Val == head.Next.Val {
				if head.Next.Next != nil {
					head.Next = head.Next.Next
				} else {
					head.Next = nil
				}
			}
		}

		head = head.Next
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	cu := head
	for cu != nil && cu.Next != nil {
		if cu.Val == cu.Next.Val {
			cu.Next = cu.Next.Next
		} else {
			cu = cu.Next
		}
	}
	return head
}

func print(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d->", head.Val)
		} else {
			fmt.Printf("%d->nil\n", head.Val)
		}
		head = head.Next
	}
}

func main() {
	a := newList(1)
	a.add(1)
	a.add(2)
	//print(a)
	deleteDuplicates2(a)
	print(a)
}
