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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	v, n := 0, 0 // v is sum, n is carry

	for {
		v, n = add(l1, l2, n)
		//fmt.Println(v, n)
		temp.Val = v

		l1 = next(l1)
		l2 = next(l2)
		// 如果两个数的下一位都为nil，则结束按位相加的运算
		if l1 == nil && l2 == nil {
			break
		}

		// 다음 작업을 위해 노드 준비 하고 링크 연결
		temp.Next = &ListNode{}
		temp = temp.Next

	}

	// n == 1은 마지막 추가 작업이 수행되었고 노드를 추가해야 함을 나타냅니다.
	if n == 1 {
		temp.Next = &ListNode{Val: n}
	}
	fmt.Println(temp)
	fmt.Println(result)
	return result
}

// next 进入l的下一位。
func next(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}

func add(n1 *ListNode, n2 *ListNode, i int) (v, n int) {
	if n1 != nil {
		v += n1.Val
	}

	if n2 != nil {
		v += n2.Val
	}

	v += i // i is carry

	if v > 9 {
		v -= 10
		n = 1
	}
	return
}

func main() {}
