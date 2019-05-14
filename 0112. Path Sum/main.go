package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	// 변경 값 반영
	sum -= root.Val

	// 탈출 조건
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	// 마지막은 재귀 ...
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum) // !!!
}

func main() {

}
