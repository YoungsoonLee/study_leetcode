package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// Go int 타입의 최소값과 최대 값
	MIN, MAX := -1<<63, 1<<63-1

	return rec(MIN, MAX, root)
}

// root.Val이 (min, max) 범위 내에 있는지 반복적으로 확인합니다.
func rec(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return min < root.Val && root.Val < max &&
		rec(min, root.Val, root.Left) &&
		rec(root.Val, max, root.Right)
}

func main() {

}
