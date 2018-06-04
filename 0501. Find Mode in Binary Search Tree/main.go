package main

import (
	"fmt"
)

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

func main() {

}
