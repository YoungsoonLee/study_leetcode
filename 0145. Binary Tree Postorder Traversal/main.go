package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append([]int{root.Val}, res...)

		dfs(root.Right)
		dfs(root.Left)
	}

	dfs(root)

	return res
}

func main() {

}
