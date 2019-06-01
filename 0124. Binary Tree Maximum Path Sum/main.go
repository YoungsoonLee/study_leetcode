package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val

	var dfs func(*TreeNode) int
	// 가능한 모든 경로의 합계 값 중 최대 값을 루트에서 시작하여 반환합니다.
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		sum := left + root.Val + right
		if maxSum < sum {
			maxSum = sum
		}
		return max(left, right) + root.Val
	}

	dfs(root)

	return maxSum
}

/*
func maxPathSum(root *TreeNode) int {
	max := 0

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left.Left == nil && root.Right.Right == nil {
			sum := root.Val + root.Left.Val + root.Right.Val
			if max < sum {
				max = sum
			}
		}

		dfs(root.Left)
		dfs(root.Right)
	}

	fmt.Println(max)
	return max
}
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
