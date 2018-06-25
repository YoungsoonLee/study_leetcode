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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k < 2 {
		return head
	}

	next, ok := needReverse(head, k)
	/*
		for next != nil {
			fmt.Println(next)
			next = next.Next
		}
	*/

	if ok {
		head, tail := reverse(head)
		fmt.Println("head, tail: ", head, tail)
		// 재귀 적
		// 미리 완성 된 k 노드의 꼬리가 조합 된 노드의 헤드를 가리킴
		tail.Next = reverseKGroup(next, k)
		return head
	}

	return head

}

func reverse(head *ListNode) (first, last *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	gotLast := false

	// node change !!!!!
	for head != nil {
		fmt.Println("head : ", head)

		temp := head.Next
		head.Next = first // !!!
		first = head      // !!!
		head = temp

		fmt.Println("temp: ", temp)
		fmt.Println("first: ", first)

		if !gotLast {
			last = first
			gotLast = true
		}
	}

	return first, last
}

// 되돌릴 필요가있는 이전 k 노드가 있는지 판별하십시오.
// 필요한 경우
// KthNode.Next = nil은 첫 번째 k 노드의 반전을 촉진하기 위해 k 및 k + 1 노드에 대해 잘립니다.
func needReverse(head *ListNode, k int) (begin *ListNode, ok bool) {
	for head != nil {
		if k == 1 {
			begin = head.Next
			// 첫 번째 k 노드와 후속 노드를 억제하여 역순으로합니다.
			head.Next = nil
			return begin, true
		}

		head = head.Next
		k--
	}

	return nil, false
}

func reverseKGroup_my(head *ListNode, k int) *ListNode {
	fast := head

	for fast != nil {
		if fast.Val == k {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
		fast = fast.Next
	}

	fast.Next = head

	return fast

}

func main() {
	l5 := &ListNode{Val: 5, Next: nil}
	l4 := &ListNode{Val: 4, Next: l5}
	l3 := &ListNode{Val: 3, Next: l4}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 1, Next: l2}
	re := reverseKGroup(l1, 3)
	for re != nil {
		fmt.Print(re.Val)
		re = re.Next
	}
}
