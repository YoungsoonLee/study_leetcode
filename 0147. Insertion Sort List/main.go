package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	// full read and save to array
	// sort
	// make it to list
	headPre := &ListNode{Next: head}
	cur := head

	for cur != nil && cur.Next != nil {
		p := cur.Next
		if cur.Val <= p.Val {
			// p는 cur 전에 삽입해야하는 요소가 아닙니다.
			cur = p
			continue
		}

		// p는 삽입해야하는 요소입니다.
		// p를 cur에서 제거한다.
		cur.Next = p.Next
		// 적절한 pre와 next 사이에 p를 삽입합니다.
		pre, next := headPre, headPre.Next
		// 올바른 장소 의미
		// pre.Val <p.Val <= next.Val
		for next.Val < p.Val {
			pre = next
			next = next.Next
		}
		// insert
		pre.Next = p
		p.Next = next
	}

	return headPre.Next
}

func main() {

}
