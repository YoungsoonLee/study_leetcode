package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	recur(root)
}

// recur는 평평하게 한 후 루트를 평평하게하고 리프 노드를 리턴한다.
func recur(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil {
		return recur(root.Right)
	}

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return recur(root.Right)
	}

	res := recur(root.Right)
	recur(root.Left).Right = root.Right

	// 왼쪽 노드를 닫는 것을 잊지 마라.
	root.Right = root.Left
	root.Left = nil

	return res
}

func main() {

}
