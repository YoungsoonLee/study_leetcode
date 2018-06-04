package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vv := make([]int, 0)
	for head != nil {
		vv = append(vv, head.Val)
		head = head.Next
	}

	i, j := 0, len(vv)-1
	for i < j {
		//for i < j/2 {	// 원소가 2개 일떄 처리가 불가능 !!!
		if vv[i] != vv[j] {
			return false
		}
		j--
		i++
	}

	return true
}

func main() {
	fmt.Println("")
}
