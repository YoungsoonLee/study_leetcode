package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := split(head)
	return merge(sortList(left), sortList(right))
}

// 중간 위치에서 목록 분할
func split(head *ListNode) (left, right *ListNode) {
	// head.Next! = nil
	// sortList가 이미 체크 아웃하는데 도움이 되었기 때문입니다.

	// 빠른 변경 속도는 두 배 빠름
	// 빠름이 끝까지 도달하면 목록의 중간에 느림
	slow, fast := head, head
	var slowPre *ListNode

	for fast != nil && fast.Next != nil {
		slowPre, slow = slow, slow.Next
		fast = fast.Next.Next
	}

	//리스트를 깨뜨린다.
	slowPre.Next = nil
	// 목록의 길이가 2 일 때 slowPre를 사용하여 길이가 1 인 두 개의 하위 목록으로 분할됩니다.
	left, right = head, slow
	return
}

// 进行合并
func merge(left, right *ListNode) *ListNode {

	cur := &ListNode{}
	headPre := cur

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}

	// 处理 left 或 right 中，剩下的节点
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}

	return headPre.Next
}

/* my solution */
/*
func sortList(head *ListNode) *ListNode {
	a := make([]int, 0)

	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}

	sort.Ints(a)

	res := &ListNode{}
	temp := res
	for _, i := range a {
		//res.Val = i
		temp.Next = &ListNode{Val: i}
		temp = temp.Next
	}

	return res.Next
}
*/

func main() {

}
