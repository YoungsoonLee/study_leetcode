/**
 * linkedlist의 reverse 함수가 매우 중요하다.
 * 꼭 다시 봐야 한다.
 */

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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	// m> = 2가되도록 headPre를 추가합니다.
	// 머리가 들어간 역의 부분은 피한다.
	// 분할 및 조합 논리를 간단하고 명확하게 만듭니다.

	headPre := &ListNode{}
	headPre.Next = head

	m++
	n++

	// m, n의 값에 따라 체인을 자른다.
	mPre, mNode, nNext := split(headPre, m, n)
	printNode(mPre)
	fmt.Println()
	printNode(mNode)
	fmt.Println()
	printNode(nNext)
	fmt.Println()

	// 중간 섹션을 뒤집습니다.
	h, e := reverse(mNode)
	/*
		printNode(h)
		fmt.Println()
		printNode(e)
		fmt.Println()
	*/

	// 체인 연결
	mPre.Next = h
	e.Next = nNext

	printNode(headPre)

	return headPre.Next
}

// 새 체인의 머리와 끝을 반환합니다.
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!
func reverse(head *ListNode) (h, e *ListNode) {
	fmt.Println(head.Val)

	if head == nil || head.Next == nil {
		return head, head
	}

	var end *ListNode

	h, end = reverse(head.Next)

	fmt.Println(head.Val)

	end.Next = head
	e = head

	/*
		printNode(h)
		fmt.Println()
		printNode(e)
		fmt.Println()
	*/

	return
}

func split(head *ListNode, m, n int) (mPre, mNode, nNext *ListNode) {
	i := 1
	for head != nil {
		if i == m-1 {
			mPre = head
			mNode = head.Next
		}
		if i == n {
			nNext = head.Next
			// head.Next = nil, 중요합니다.
			//하지 않으면 역방향이 잘못됩니다.
			head.Next = nil
			break
		}
		head = head.Next
		i++
	}

	return
}

func printNode(node *ListNode) {
	for node != nil {
		fmt.Print("-> ", node.Val)
		node = node.Next
	}
	//fmt.Print("-> ", node.Val)
}

// !!!!!
func makeAtoL(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}

	res := &ListNode{Val: a[0]}

	temp := res //shalow copy !!!

	for i := 1; i < len(a); i++ {
		temp.Next = &ListNode{Val: a[i]}
		temp = temp.Next
	}

	return res
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	m := 2
	n := 5
	reverseBetween(makeAtoL(a), m, n)
}
