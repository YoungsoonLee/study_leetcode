/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
 *
 * https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (59.68%)
 * Likes:    887
 * Dislikes: 129
 * Total Accepted:    89.4K
 * Total Submissions: 149.6K
 * Testcase Example:  '[3,9,20,15,7]'
 *
 * Given a non-empty binary tree, return the average value of the nodes on each
 * level in the form of an array.
 * 
 * Example 1:
 * 
 * Input:
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * Output: [3, 14.5, 11]
 * Explanation:
 * The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on
 * level 2 is 11. Hence return [3, 14.5, 11].
 * 
 * 
 * 
 * Note:
 * 
 * The range of node's value is in the range of 32-bit signed integer.
 * 
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	nodes := make([]*TreeNode,1, 1024)
	nodes[0] = root

	fmt.Println(nodes)

	for len(nodes) > 0 {
		n := len(nodes)

		sum := 0

		for i:=0; i<n;i++ {
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

