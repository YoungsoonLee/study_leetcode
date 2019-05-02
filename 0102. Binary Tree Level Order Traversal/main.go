package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)

	fmt.Println(res)
	fmt.Println(len(res))

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 새로운 레벨이 나타났습니다.
		if level > len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
	/*
		result := make([][]int, 0)

		// save root
		result = append(result, []int{root.Val})

		for root.Left != nil && root.Right != nil {
			recursive(root.Left, &result)
			recursive(root.Right, &result)
			if root.Left != nil {
				root = root.Left
			}
			if root.Right != nil {
				root = root.Right
			}
		}

		fmt.Println(result)
		return result
	*/
}

func recursive(node *TreeNode, result *[][]int) {
	*result = append(*result, []int{node.Val})
}

func main() {

}
