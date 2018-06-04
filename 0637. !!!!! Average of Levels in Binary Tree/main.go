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

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0, 128)

	nodes := make([]*TreeNode, 1, 1024)
	nodes[0] = root

	//fmt.Println(nodes)

	for len(nodes) > 0 {
		n := len(nodes)

		sum := 0

		for i := 0; i < n; i++ {
			//fmt.Println(n, i)
			sum += nodes[i].Val
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}

			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}

		res = append(res, float64(sum)/float64(n))
		nodes = nodes[n:]
	}

	return res
}

func main() {}
