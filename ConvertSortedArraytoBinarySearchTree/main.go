/**
 * 1. 정렬 및 발란스 2진 트리 되어 있으므로 중간이 항상 root 이다.
 * 2. 루트 보다 작으면 왼쪽, 크면 오른쪽 ... recursive
 * */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// my solution - fail
/*
func sortedArrayToBST(nums []int) *TreeNode {
	root := len(nums) / 2
	t := &TreeNode{Val: nums[root], Left: nil, Right: nil} // init tree

	nums = append(nums[0:root], nums[root+1:]...)
	//fmt.Println(nums)
	//fmt.Println("t: ", t)
	for _, n := range nums {
		//fmt.Println("n: ", n)
		bst(t, n)
		nums = nums[1:]
		//fmt.Println(nums)
	}
	return t
}

func bst(t *TreeNode, val int) {

	//for t.Left != nil || t.Right != nil {
	fmt.Println("t: ", t)
	if val < t.Val {
		if t.Left == nil {
			// insert
			t.Left = &TreeNode{Val: val, Left: nil, Right: nil}

			// compare & change

			if t.Val < t.Left.Val {
				//change val
				temp := t.Left.Val
				t.Left.Val = t.Val
				t.Val = temp
			}
		} else {
			t = t.Left
			bst(t, val)
		}

	} else {
		if t.Right == nil {
			// insert
			t.Right = &TreeNode{Val: val, Left: nil, Right: nil}

			// compare & change
			if t.Val > t.Right.Val {
				//change val
				temp := t.Right.Val
				t.Right.Val = t.Val
				t.Val = temp
			}

		} else {
			t = t.Right
			bst(t, val)
		}
	}
	//}
}
*/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

func main() {
	a := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(a))
}
