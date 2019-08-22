package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func findMode(root *TreeNode) []int {
	r := map[int]int{}
	search(root, r)
	fmt.Println("r: ", r)

	max := -1
	res := []int{}
	for n, v := range r {
		if max <= v {
			if max < v {
				max = v
				res = res[0:0]
			}
			res = append(res, n)
		}
	}

	fmt.Println(res)
	return res
}

// dfs
func search(root *TreeNode, rec map[int]int) {
	if root == nil {
		return
	}

	rec[root.Val]++

	search(root.Left, rec)
	search(root.Right, rec)
}
*/

func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	res := make([]int, 0)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if _, ok := m[root.Val]; ok {
			m[root.Val]++
			if m[root.Val] >= 2 {
				res = append(res, m[root.Val])
			}
		} else {
			m[root.Val]++
		}

		if root.Left != nil {
			dfs(root.Left)
		}

		if root.Right != nil {
			dfs(root.Right)
		}

	}

	dfs(root)

	return res

}

func main() {

}
