package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	count := 0

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		count++

		dfs(root.Left)
		dfs(root.Right)

	}

	dfs(root)

	return count
}

func main() {

}
