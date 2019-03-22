package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// <x 노드들의 체인을 저장한다.
	lessHead := &ListNode{}
	// 노드 저장 체인 >= x
	noLessHead := &ListNode{}

	// Head.Next는 실제 머리입니다.
	// for 루프의 논리에 대해 수행됩니다.
	lessEnd := lessHead
	noLessEnd := noLessHead

	for head != nil {
		if head.Val < x {
			lessEnd.Next = head
			lessEnd = lessEnd.Next
		} else {
			noLessEnd.Next = head
			noLessEnd = noLessEnd.Next
		}
		head = head.Next
	}

	/*
		fmt.Println("---- head ----")
		for head != nil {
			fmt.Printf("%d -> ", head.Val)
			head = head.Next
		}
		fmt.Println()
		fmt.Println("---- lessHead ----")
		for lessHead != nil {
			fmt.Printf("%d -> ", lessHead.Val)
			lessHead = lessHead.Next
		}
		fmt.Println()
		fmt.Println("---- noLessHead ----")
		for noLessHead != nil {
			fmt.Printf("%d -> ", noLessHead.Val)
			noLessHead = noLessHead.Next
		}
		fmt.Println()
		fmt.Println("---- lessEnd ----")
		for lessEnd != nil {
			fmt.Printf("%d -> ", lessEnd.Val)
			lessEnd = lessEnd.Next
		}
		fmt.Println()
		fmt.Println("---- noLessEnd ----")
		for noLessEnd != nil {
			fmt.Printf("%d -> ", noLessEnd.Val)
			noLessEnd = noLessEnd.Next
		}
	*/

	// 두 부분을 끝에서 끝까지 연결합니다.
	lessEnd.Next = noLessHead.Next
	// 닫힌 노트는 noLessEnd입니다.
	noLessEnd.Next = nil

	head = lessHead.Next

	return head
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

func main() {
	n := []int{1, 4, 3, 2, 5, 2}
	x := 3
	result := partition(s2l(n), x)

	for result != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}
}
