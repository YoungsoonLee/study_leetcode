package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	temp := head

	for temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	//fmt.Println(head)

	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}

	fmt.Println()
	fmt.Println(temp.Val)

	for temp != nil {
		fmt.Print(temp.Val, " -> ")
		temp = temp.Next
	}
	return head

}

func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res // TODO: WHY!!!!!!!!!!!!!!!!!!! for sending head ???

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

func main() {
	n := []int{1, 1, 2}
	deleteDuplicates(s2l(n))
}
