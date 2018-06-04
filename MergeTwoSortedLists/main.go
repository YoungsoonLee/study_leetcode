package main

import (
	"fmt"
	"sort"
)

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

func newList(value int) *ListNode {
	return &ListNode{Val: value, Next: nil}
}

func (l *ListNode) Add(value int) {
	if l == nil {
		l.Val = value
		l.Next = nil
	} else {
		// TODO: need to sorted insert
		// first assume: input already sorted data
		for l.Next != nil {
			l = l.Next
		}
		l.Next = &ListNode{Val: value, Next: nil}
	}
}

func (l *ListNode) Print() {
	for l != nil {
		if l.Next == nil {
			fmt.Printf("%d -> nil \n", l.Val)
		} else {
			fmt.Printf("%d -> ", l.Val)
		}

		l = l.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	re := make([]int, 0)
	for l1 != nil {
		re = append(re, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		re = append(re, l2.Val)
		l2 = l2.Next
	}

	//fmt.Println(re)
	sort.Sort(sort.IntSlice(re))

	reList := newList(re[0])
	for _, v := range re[1:] {
		reList.Add(v)
	}

	return reList
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	// best logic ~!!!!!
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}

func main() {
	l1 := newList(1)
	l1.Add(2)
	l1.Add(4)
	//l1.Print()

	l2 := newList(1)
	l2.Add(3)
	l2.Add(4)
	//l2.Print()

	rs := mergeTwoLists2(l1, l2)
	rs.Print()

}
