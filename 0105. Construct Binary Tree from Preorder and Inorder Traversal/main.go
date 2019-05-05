package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	res := &TreeNode{Val: preorder[0]}

	if len(inorder) == 1 {
		return res
	}

	idx := indexOf(res.Val, inorder)
	fmt.Println(idx)

	res.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	res.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return res

	/*
		t := TreeNode{Val: preorder[0], Left: nil, Right: nil} // init

		var dfs func([]int, *TreeNode)
		dfs = func(pre []int, root *TreeNode) {
			if len(pre) == 0 {
				return
			}

			pre = pre[1:]

			if root.Val > pre[0] {
				dfs(pre, root.Left)
			} else {
				dfs(pre, root.Right)
			}
		}

		dfs(preorder[1:], &t)
		return &t
	*/
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}

func main() {
	p := []int{3, 9, 20, 15, 7}
	i := []int{9, 3, 15, 20, 7}

	buildTree(p, i)
}
