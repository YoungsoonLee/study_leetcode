package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const intMax = 1<<63 - 1

func findSecondMinimumValue(root *TreeNode) int {
	fmt.Println(intMax)
	res := intMax

	helper(root, root.Val, &res)

	if res == intMax {
		return -1
	}

	return res

}

func helper(root *TreeNode, lo int, hi *int) {
	if root == nil {
		return
	}

	if lo < root.Val && root.Val < *hi {
		*hi = root.Val
	}

	helper(root.Left, lo, hi)
	helper(root.Right, lo, hi)
}

func main() {}
