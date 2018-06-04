package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* my solution - failed
func hasPathSum(root *TreeNode, sum int) bool {
	var res int

	if root == nil {
		return false
	}

	if root.Val == sum {
		return true
	} else {
		res = recur(root, res, sum)
	}

	return res == sum
}

func recur(t *TreeNode, s, sum int) int {
	s += t.Val
	if s == sum {
		return s
	}
	if t.Left != nil {
		recur(t.Left, s, sum)
	} else if t.Right != nil {
		recur(t.Right, s, sum)
	} else {
		return s
	}
	return s
}
*/

// DFS recursive
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val

	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum) // !!!
	//hasPathSum(root.Left, sum)
	//hasPathSum(root.Right, sum)
	//return sum
}

func main() {

}
