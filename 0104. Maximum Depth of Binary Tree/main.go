package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	//max := 0
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
func arrToTree(arr []int) *TreeNode {
	t := &TreeNode{Val: arr[0], Left: nil, Right: nil}

	for _, i := range arr[1:] {
		// TODO:
	}

	return t
}
*/

func main() {

}
