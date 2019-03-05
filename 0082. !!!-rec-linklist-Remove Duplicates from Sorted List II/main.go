package main

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

func deleteDuplicates(head *ListNode) *ListNode {
	// 길이가 <= 1 인 목록을 직접 반환 할 수 있습니다.
	if head == nil || head.Next == nil {
		return head
	}

	// 헤드가 반복되거나, 헤드가 삭제됩니다.
	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		// 재귀
		return deleteDuplicates(head.Next)
	}

	//Head가 반복되지 않고, 머리 뒤로 노드를 재귀 적으로 처리합니다.
	head.Next = deleteDuplicates(head.Next)
	return head
}

func main() {
	n := []int{1, 1, 1, 2, 3}
	deleteDuplicates(s2l(n))
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
