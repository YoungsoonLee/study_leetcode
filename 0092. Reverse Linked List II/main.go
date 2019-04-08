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

	temp := head
	a := make([]int, 0)

	for i := 1; i <= n; i++ {
		if i >= m && i <= n {
			a = append(a, temp.Val)
		}
		temp = temp.Next
	}

	fmt.Println(a)

	//make

	return temp
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
	a := []int{1, 2, 3, 4, 5}
	m := 2
	n := 4
	reverseBetween(makeAtoL(a), m, n)
}
