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

/*
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	m := make(map[int]float64)

	//
	// res = append(res, float64(root.Val))

	var dfs func(*TreeNode, int)
	i := 0
	dfs = func(root *TreeNode, i int) {
		if _, ok := m[i]; ok {
			m[i] = float64((int(m[i]) / root.Val) / 2)
		} else {
			m[i] = float64(root.Val)
		}
		i++
		dfs(root.Left, i)
		dfs(root.Right, i)
	}
	dfs(root, i)

	for _, v := range m {
		res = append(res, v)
	}

	return res
}
*/

func main() {}
