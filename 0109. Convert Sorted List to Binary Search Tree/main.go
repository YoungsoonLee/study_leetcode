package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return transMidToRoot(head, nil)
}

func transMidToRoot(begin, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}

	if begin.Next == end {
		return &TreeNode{Val: begin.Val}
	}

	fast, slow := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow

	return &TreeNode{
		Val:   mid.Val,
		Left:  transMidToRoot(begin, mid),
		Right: transMidToRoot(mid.Next, end),
	}

}

/*
func sortedListToBST(head *ListNode) *TreeNode {
	// treavse list one time
	// save listnode to array
	temp := []int{}
	for head != nil {
		temp = append(temp, head.Val)
		if head.Next != nil {
			head = head.Next
		}
	}

	return sBST(temp)

}

func sBST(arr []int) *TreeNode {
	mid := len(arr) / 2
	return &TreeNode{
		Val:   arr[mid],
		Left:  sBST(arr[:mid]),
		Right: sBST(arr[mid+1:]),
	}
}
*/

func main() {

}
