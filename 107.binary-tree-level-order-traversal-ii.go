/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (49.45%)
 * Likes:    1131
 * Dislikes: 203
 * Total Accepted:    292.5K
 * Total Submissions: 576.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 * 
 * 
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 
 * return its bottom-up level order traversal as:
 * 
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if level >= len((*res)) {
		// 새로운 레벨이 나타남
		(*res) = append([][]int{[]int{}}, (*res)...)
	}

	n := len((*res))
	fmt.Println(level, n)
	
	(*res)[n-level-1] = append((*res)[n-level-1], root.Val)

	dfs(root.Left, level+1, res)
	dfs(root.Right, level+1, res)

}
// @lc code=end

