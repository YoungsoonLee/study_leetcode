package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func my_isPalindrome(head *ListNode) bool {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	if len(res) == 0 {
		return true
	}

	fmt.Print(res)

	i, j := 0, len(res)-1

	for i <= j {
		fmt.Println(i, j, res[i], res[j])

		if res[i] != res[j] {
			return false
		}
		i++
		j--
	}

	return true
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
