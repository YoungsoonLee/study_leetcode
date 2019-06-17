package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	cur := head
	size := 0
	for cur != nil {
		cur = cur.Next
		size++
	}

	// 크기가 이상하다. cur가리스트의 중간 노드를 가리킨다.
	// 크기는 짝수이며, cur은 목록의 처음 절반에있는 마지막 노드를 가리 킵니다.
	cur = head
	for i := 0; i < (size-1)/2; i++ {
		cur = cur.Next
	}

	// head -> 1 -> 2 -> 3 -> 4 -> 5 -> 6
	//                   ^
	//                   |
	//                  cur

	// 역행렬 뒤에있는 목록
	next := cur.Next
	for next != nil {
		temp := next.Next
		next.Next = cur
		cur = next
		next = temp
	}
	end := cur

	// head -> 1 -> 2 -> 3 <-> 4 <- 5 <- 6 <- end
	// 양쪽 끝에서 시작하여 체인을 통합합니다.
	for head != end {
		hNext := head.Next
		eNext := end.Next
		head.Next = end
		end.Next = hNext
		head = hNext
		end = eNext
	}

	// 링 목록을 피하기 위해 목록 닫기
	end.Next = nil
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func main() {
	a := []int{1, 2, 3, 4}
	reorderList(Ints2List(a))
}
