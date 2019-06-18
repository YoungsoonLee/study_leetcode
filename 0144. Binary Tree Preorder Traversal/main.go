package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)

		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	return res
}

func main() {

}
