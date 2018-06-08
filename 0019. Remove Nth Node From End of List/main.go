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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	d, headIsNthFromEnd := getDaddy(head, n)

	if headIsNthFromEnd {
		// 헤드 노드를 삭제합니다.
		return head.Next
	}

	/*
		for d != nil {
			fmt.Println(d.Val)
			d = d.Next
		}
	*/

	d.Next = d.Next.Next
	return head

}

func getDaddy(head *ListNode, n int) (daddy *ListNode, headIsNthFromEnd bool) {
	daddy = head // shallow copy

	for head != nil {
		if n < 0 {
			daddy = daddy.Next
		}

		n--
		head = head.Next
	}

	fmt.Println(n, head, daddy)
	// n == 0, 체인의 길이가 n과 동일 함을 나타냅니다.
	headIsNthFromEnd = n == 0
	return
}

func removeNthFromEnd_my(head *ListNode, n int) *ListNode {
	m := make(map[int]int)

	i := 1
	temp := head

	for temp.Next != nil {
		m[temp.Val] = i
		i++
		temp = temp.Next
	}

	// last
	m[temp.Val] = i
	fmt.Println(m, len(m), m[len(m)-n])

	for i := 0; i < len(m); i++ {
		if i == m[len(m)-n-1] {
			head.Next = head.Next.Next
		}
		head = head.Next
	}

	fmt.Println(head)

	return head
}

func main() {
	t4 := &ListNode{Val: 5, Next: nil}
	t3 := &ListNode{Val: 4, Next: t4}
	t2 := &ListNode{Val: 3, Next: t3}
	t1 := &ListNode{Val: 2, Next: t2}
	t0 := &ListNode{Val: 1, Next: t1}

	fmt.Println(removeNthFromEnd(t0, 2))

}
